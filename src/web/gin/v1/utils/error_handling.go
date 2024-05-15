package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gym-management/src/lib/primitives/application_specific"
	"net/http"
)

type HttpErrorResponse struct {
	Method        string                 `json:"method"`
	Path          string                 `json:"path"`
	Status        int                    `json:"status"`
	CorrelationId string                 `json:"correlationId"`
	Error         HttpErrorResponseError `json:"error"`
}

type HttpErrorResponseError struct {
	Code    string            `json:"code"`
	Message string            `json:"message"`
	Payload map[string]string `json:"payload"`
}

func GlobalErrorHandler(c *gin.Context, err any) {
	switch e := err.(type) {
	case error:
		HandleError(c, e)
	case string:
		HandleError(c, errors.New(e))
	default:
		HandleError(c, errors.New("unknown error"))
	}
}

func HandleError(c *gin.Context, err error) {
	switch v := err.(type) {
	case *application_specific.ApplicationException:
		handleApplicationException(c, v)
	default:
		handleGenericError(c, v)
	}
}

func handleGenericError(c *gin.Context, err error) {
	session := ExtractSession(c)

	c.JSON(http.StatusInternalServerError, HttpErrorResponse{
		Method:        c.Request.Method,
		Path:          c.Request.RequestURI,
		Status:        http.StatusInternalServerError,
		CorrelationId: session.CorrelationId,
		Error: HttpErrorResponseError{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
			Payload: nil,
		},
	})
	c.Abort()
}

func handleApplicationException(c *gin.Context, err *application_specific.ApplicationException) {
	session := ExtractSession(c)

	var status int

	if application_specific.IsValidationException(err) {
		status = http.StatusBadRequest
	}

	if application_specific.IsAuthenticationException(err) {
		status = http.StatusUnauthorized
	}

	if application_specific.IsForbiddenException(err) {
		status = http.StatusForbidden
	}

	if application_specific.IsNotFoundException(err) {
		status = http.StatusNotFound
	}

	if application_specific.IsUnknownException(err) {
		status = http.StatusInternalServerError
	}

	c.JSON(status, HttpErrorResponse{
		Method:        c.Request.Method,
		Path:          c.Request.RequestURI,
		Status:        status,
		CorrelationId: session.CorrelationId,
		Error: HttpErrorResponseError{
			Code:    err.Code,
			Message: err.Message,
			Payload: err.Payload,
		},
	})
	c.Abort()
}

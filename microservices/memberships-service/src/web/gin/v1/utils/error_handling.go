package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gym-management-memberships/src/lib/primitives/application_specific"
	"net/http"
)

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
	var correlationId string

	exists := CheckSession(c)
	if exists {
		correlationId = ExtractSession(c).CorrelationId
	}

	c.JSON(http.StatusInternalServerError, HttpErrorResponse{
		Method:        c.Request.Method,
		Path:          c.Request.RequestURI,
		Status:        http.StatusInternalServerError,
		CorrelationId: correlationId,
		Error: HttpErrorResponseError{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
			Payload: nil,
		},
	})
	c.Abort()
}

func handleApplicationException(c *gin.Context, err *application_specific.ApplicationException) {
	var correlationId string

	exists := CheckSession(c)
	if exists {
		correlationId = ExtractSession(c).CorrelationId
	}

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
		CorrelationId: correlationId,
		Error: HttpErrorResponseError{
			Code:    err.Code,
			Message: err.Message,
			Payload: err.Payload,
		},
	})
	c.Abort()
}

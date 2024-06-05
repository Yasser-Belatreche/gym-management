package utils

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"gym-management-gyms/src/lib/primitives/application_specific"
)

func NewInvalidRequestError(err error) *application_specific.ApplicationException {
	var errors = make(map[string]string)

	switch ve := err.(type) {
	case validator.ValidationErrors:
		for _, fe := range ve {
			field := fe.Field()
			tag := fe.Tag()
			errors[field] = formatErrorMessage(field, tag)
		}
	case *json.UnmarshalTypeError:
		errors[ve.Field] = fmt.Sprintf("Invalid type for field '%s', expected '%s'", ve.Field, ve.Type.String())
	case *json.SyntaxError:
		errors["json"] = fmt.Sprintf("JSON syntax error at byte offset %d", ve.Offset)
	default:
		if ve.Error() == "json: cannot unmarshal number into Go value of type string" {
			errors["json"] = "Invalid JSON format: incorrect data types"
		} else {
			errors["json"] = ve.Error()
		}
	}

	return application_specific.NewValidationException("INVALID_REQUEST", "invalid request body", errors)
}

func formatErrorMessage(field, tag string) string {
	var message string
	switch tag {
	case "required":
		message = fmt.Sprintf("The %s field is required", field)
	case "email":
		message = fmt.Sprintf("The %s field must be a valid email address", field)
	case "max":
		message = fmt.Sprintf("The %s field must be less than the maximum value", field)
	case "min":
		message = fmt.Sprintf("The %s field must be more than the minimum value", field)
	default:
		message = fmt.Sprintf("The %s field is invalid", field)
	}
	return message
}

func NewRouteNotFoundError() *application_specific.ApplicationException {
	return application_specific.NewNotFoundException("ROUTE_NOT_FOUND", "Route not found", nil)
}

func NewNoApiSecretError() *application_specific.ApplicationException {
	return application_specific.NewDeveloperException("NO_API_SECRET", "No api secret provided")
}

func NewWrongApiSecretError() *application_specific.ApplicationException {
	return application_specific.NewDeveloperException("WRONG_API_SECRET", "Wrong api secret")
}

func NewNoSessionError() *application_specific.ApplicationException {
	return application_specific.NewDeveloperException("NO_SESSION", "No session provided")
}

func NewInvalidSessionError() *application_specific.ApplicationException {
	return application_specific.NewDeveloperException("INVALID_SESSION", "Session should be an encoded json string in base64")
}

func NewNoUserSessionError() *application_specific.ApplicationException {
	return application_specific.NewAuthenticationException("NO_USER_SESSION", "No user session provided", nil)
}

package application_specific

const (
	validationExceptionType     = "ValidationException"
	notFoundExceptionType       = "NotFoundException"
	authenticationExceptionType = "AuthenticationException"
	forbiddenExceptionType      = "ForbiddenException"
	unknownExceptionType        = "UnknownException"
	developerExceptionType      = "DeveloperException"
)

type ApplicationException struct {
	Code      string
	Message   string
	Payload   map[string]interface{}
	exception string
}

func (exception *ApplicationException) Error() string {
	return exception.Message
}

func NewValidationException(code string, message string, payload map[string]interface{}) *ApplicationException {
	return &ApplicationException{
		Code:      code,
		Message:   message,
		Payload:   payload,
		exception: validationExceptionType,
	}
}

func NewNotFoundException(code string, message string, payload map[string]interface{}) *ApplicationException {
	return &ApplicationException{
		Code:      code,
		Message:   message,
		Payload:   payload,
		exception: notFoundExceptionType,
	}
}

func NewAuthenticationException(code string, message string, payload map[string]interface{}) *ApplicationException {
	return &ApplicationException{
		Code:      code,
		Message:   message,
		Payload:   payload,
		exception: authenticationExceptionType,
	}
}

func NewForbiddenException(code string, message string, payload map[string]interface{}) *ApplicationException {
	return &ApplicationException{
		Code:      code,
		Message:   message,
		Payload:   payload,
		exception: forbiddenExceptionType,
	}
}

func NewUnknownException(code string, message string, payload map[string]interface{}) *ApplicationException {
	return &ApplicationException{
		Code:      code,
		Message:   message,
		Payload:   payload,
		exception: unknownExceptionType,
	}
}

func NewDeveloperException(code string, message string) *ApplicationException {
	return &ApplicationException{
		Code:      "DEVELOPER." + code,
		Message:   message,
		Payload:   nil,
		exception: developerExceptionType,
	}
}

func IsValidationException(err interface{}) bool {
	switch e := err.(type) {
	case *ApplicationException:
		return e.exception == validationExceptionType
	case ApplicationException:
		return e.exception == validationExceptionType
	default:
		return false
	}
}

func IsNotFoundException(err interface{}) bool {
	switch e := err.(type) {
	case *ApplicationException:
		return e.exception == notFoundExceptionType
	case ApplicationException:
		return e.exception == notFoundExceptionType
	default:
		return false
	}
}

func IsAuthenticationException(err interface{}) bool {
	switch e := err.(type) {
	case *ApplicationException:
		return e.exception == authenticationExceptionType
	case ApplicationException:
		return e.exception == authenticationExceptionType
	default:
		return false
	}
}

func IsForbiddenException(err interface{}) bool {
	switch e := err.(type) {
	case *ApplicationException:
		return e.exception == forbiddenExceptionType
	case ApplicationException:
		return e.exception == forbiddenExceptionType
	default:
		return false
	}
}

func IsUnknownException(err interface{}) bool {
	switch e := err.(type) {
	case *ApplicationException:
		return e.exception == unknownExceptionType || e.exception == developerExceptionType
	case ApplicationException:
		return e.exception == unknownExceptionType || e.exception == developerExceptionType
	default:
		return false
	}
}

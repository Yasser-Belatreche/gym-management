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
	code      string
	message   string
	payload   map[string]string
	exception string
}

func (exception *ApplicationException) Error() string {
	return exception.message
}

func NewValidationException(code string, message string, payload map[string]string) *ApplicationException {
	return &ApplicationException{
		code:      code,
		message:   message,
		payload:   payload,
		exception: validationExceptionType,
	}
}

func NewNotFoundException(code string, message string, payload map[string]string) *ApplicationException {
	return &ApplicationException{
		code:      code,
		message:   message,
		payload:   payload,
		exception: notFoundExceptionType,
	}
}

func NewAuthenticationException(code string, message string, payload map[string]string) *ApplicationException {
	return &ApplicationException{
		code:      code,
		message:   message,
		payload:   payload,
		exception: authenticationExceptionType,
	}
}

func NewForbiddenException(code string, message string, payload map[string]string) *ApplicationException {
	return &ApplicationException{
		code:      code,
		message:   message,
		payload:   payload,
		exception: forbiddenExceptionType,
	}
}

func NewUnknownException(code string, message string, payload map[string]string) *ApplicationException {
	return &ApplicationException{
		code:      code,
		message:   message,
		payload:   payload,
		exception: unknownExceptionType,
	}
}

func NewDeveloperException(code string, message string) *ApplicationException {
	return &ApplicationException{
		code:      code,
		message:   message,
		payload:   nil,
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

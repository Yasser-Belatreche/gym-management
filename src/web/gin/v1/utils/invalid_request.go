package utils

import "gym-management/src/lib/primitives/application_specific"

func NewInvalidRequestError(err error) *application_specific.ApplicationException {
	return application_specific.NewValidationException("INVALID_REQUEST", err.Error(), nil)
}

func NewRouteNotFoundError() *application_specific.ApplicationException {
	return application_specific.NewNotFoundException("ROUTE_NOT_FOUND", "Route not found", nil)
}

func NewNoTokenError() *application_specific.ApplicationException {
	return application_specific.NewAuthenticationException("NO_TOKEN", "No token provided", nil)
}

package application_specific

func AssertUserRole(session *UserSession, roles ...string) *ApplicationException {
	if session.RoleIsOneOf(roles...) {
		return nil
	}

	return NewForbiddenException("OPERATION_NOT_ALLOWED", "You are not allowed to do this operation", map[string]string{
		"id": session.UserId(),
	})
}

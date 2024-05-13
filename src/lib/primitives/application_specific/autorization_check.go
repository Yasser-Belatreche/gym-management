package application_specific

func AssertUserRole(session *UserSession, roles ...string) *ApplicationException {
	if session.RoleIsOneOf(roles...) {
		return nil
	}

	allowedRoles := ""
	for i, role := range roles {
		allowedRoles += role
		if i != len(roles)-1 {
			allowedRoles += ", "
		}
	}

	return NewForbiddenException("OPERATION_NOT_ALLOWED", "You are not allowed to do this operation", map[string]string{
		"userId":       session.User.Id,
		"userRole":     session.User.Role,
		"allowedRoles": allowedRoles,
	})
}

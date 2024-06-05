package domain

import "gym-management-auth/src/lib/primitives/application_specific"

type Username string

func UsernameFrom(str string) (Username, *application_specific.ApplicationException) {
	if application_specific.IsEmailValid(str) {
		email, _ := application_specific.NewEmail(str)

		return UsernameFromEmail(email), nil
	}

	return "", application_specific.NewValidationException("AUTH.USERNAME.INVALID", str+" is not a valid username", map[string]string{
		"username": str,
	})
}

func UsernameFromEmail(email application_specific.Email) Username {
	return Username(email.Value)
}

func (u Username) Value() string {
	return string(u)
}

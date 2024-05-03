package application_specific

import (
	"regexp"
	"strings"
)

type Email struct {
	Value string
}

const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

func NewEmail(str string) (Email, *ApplicationException) {
	str = strings.ToLower(strings.TrimSpace(str))

	if !isEmailValid(str) {
		return Email{}, NewValidationException("EMAIL.INVALID_EMAIL", str+" is not a valid email", map[string]string{
			"email": str,
		})
	}

	return Email{Value: str}, nil
}

func (e Email) Equals(another Email) bool {
	return e.Value == another.Value
}

func isEmailValid(email string) bool {
	matched, _ := regexp.MatchString(emailRegex, email)

	return matched
}

package domain

import (
	"golang.org/x/crypto/bcrypt"
	"gym-management/src/lib/primitives/application_specific"
	"strings"
)

type Password string

func PasswordFromPlain(str string) (Password, *application_specific.ApplicationException) {
	hash, err := bcrypt.GenerateFromPassword([]byte(str), 14)
	if err != nil {
		return "", application_specific.NewUnknownException("AUTH.PASSWORD.HASHING_FAILED", "Password hashing failed", map[string]string{
			"password": str,
		})
	}

	return Password(hash), nil
}

func PasswordFromEncrypted(str string) (Password, *application_specific.ApplicationException) {
	if !isEncrypted(str) {
		return "", application_specific.NewValidationException("AUTH.PASSWORD.INVALID", "Password is not encrypted", map[string]string{
			"password": str,
		})
	}

	return Password(str), nil
}

func (p Password) Value() string {
	return string(p)
}

func (p Password) Equals(another string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(p.Value()), []byte(another))

	return err == nil
}

func isEncrypted(str string) bool {
	return strings.HasPrefix(str, "$2a")
}

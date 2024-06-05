package application_specific

import "testing"

func TestIsValidationException(t *testing.T) {
	err := NewValidationException("Code", "Message", nil)

	if !IsValidationException(err) {
		t.Errorf("Expected true, got false")
	}

	notValidation := NewNotFoundException("Code", "Message", nil)

	if IsValidationException(notValidation) {
		t.Errorf("Expected false, got true")
	}
}

func TestIsNotFoundException(t *testing.T) {
	err := NewNotFoundException("Code", "Message", nil)

	if !IsNotFoundException(err) {
		t.Errorf("Expected true, got false")
	}

	notNotFound := NewAuthenticationException("Code", "Message", nil)

	if IsNotFoundException(notNotFound) {
		t.Errorf("Expected false, got true")
	}
}

func TestIsAuthenticationException(t *testing.T) {
	err := NewAuthenticationException("Code", "Message", nil)

	if !IsAuthenticationException(err) {
		t.Errorf("Expected true, got false")
	}

	notAuth := NewForbiddenException("Code", "Message", nil)

	if IsAuthenticationException(notAuth) {
		t.Errorf("Expected false, got true")
	}
}

func TestIsForbiddenException(t *testing.T) {
	err := NewForbiddenException("Code", "Message", nil)

	if !IsForbiddenException(err) {
		t.Errorf("Expected true, got false")
	}

	notForbidden := NewUnknownException("Code", "Message", nil)

	if IsForbiddenException(notForbidden) {
		t.Errorf("Expected false, got true")

	}
}

func TestIsUnknownException(t *testing.T) {
	err := NewUnknownException("Code", "Message", nil)

	if !IsUnknownException(err) {
		t.Errorf("Expected true, got false")
	}

	notUnknown := NewValidationException("Code", "Message", nil)

	if IsUnknownException(notUnknown) {
		t.Errorf("Expected false, got true")
	}
}

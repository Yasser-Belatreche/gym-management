package domain

import exceptions "gym-management/src/lib/primitives/application_specific"

type Gender string

const (
	Male   = "male"
	Female = "female"
)

func NewGender(str string) (Gender, *exceptions.ApplicationException) {
	if str == Male {
		return Gender(str), nil
	}

	if str == Female {
		return Gender(str), nil
	}

	return "", exceptions.NewValidationException("GYMS.CUSTOMERS.INVALID_GENDER", str+" is not a valid gender", map[string]string{
		"gender": str,
	})
}

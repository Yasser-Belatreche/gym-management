package domain

import exceptions "gym-management/src/lib/primitives/application_specific"

type Gender string

const (
	GenderMale   = "male"
	GenderFemale = "female"
)

func GenderFrom(str string) (Gender, *exceptions.ApplicationException) {
	if str == GenderMale {
		return Gender(str), nil
	}

	if str == GenderFemale {
		return Gender(str), nil
	}

	return "", exceptions.NewValidationException("GYMS.CUSTOMERS.INVALID_GENDER", str+" is not a valid gender", map[string]string{
		"gender": str,
	})
}

func (g Gender) Value() string {
	return string(g)
}

package generic

import "testing"

func TestGenerateRandomStringLength(t *testing.T) {
	length := 30
	str := GenerateRandomString(length)
	if len(str) != length {
		t.Errorf("Expected length %d, got %d", length, len(str))
	}
}

func TestGenerateRandomStringRandomness(t *testing.T) {
	length := 10
	str1 := GenerateRandomString(length)
	str2 := GenerateRandomString(length)
	str3 := GenerateRandomString(length)

	if str1 == str2 && str2 == str3 {
		t.Errorf("Generated strings are not random")
	}
}

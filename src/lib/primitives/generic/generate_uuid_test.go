package generic

import "testing"

func TestGenerateUUIDShouldReturnUniqueIds(t *testing.T) {
	if GenerateULID() == GenerateULID() {
		t.Errorf("Expected different ids, got the same")
	}
}

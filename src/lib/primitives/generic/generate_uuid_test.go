package generic

import "testing"

func TestGenerateUUIDShouldReturnUniqueIds(t *testing.T) {
	if GenerateUUID() == GenerateUUID() {
		t.Errorf("Expected different ids, got the same")
	}
}

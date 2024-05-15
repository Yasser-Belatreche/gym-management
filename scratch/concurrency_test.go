package scratch

import (
	"testing"
)

func TestSearchWord(t *testing.T) {
	t.Skip("skipping test")
	fileContents := map[string]string{
		"file1.txt": "apple banana apple banana",
		"file2.txt": "apple orange banana apple",
		"file3.txt": "orange banana apple orange",
	}

	info1, err1 := SearchWord("apple", fileContents)
	if err1 != nil {
		t.Errorf("Test case 1 failed: %v", err1)
	} else if info1.TF != 5 || info1.DF != 3 {
		t.Errorf("Test case 1 failed: expected TF=5, DF=3, got TF=%d, DF=%d", info1.TF, info1.DF)
	} else if info1.SearchCount != 1 {
		t.Errorf("Test case 1 failed: expected SearchCount=1, got SearchCount=%d", info1.SearchCount)
	}

	info2, err2 := SearchWord("banana", fileContents)
	if err2 != nil {
		t.Errorf("Test case 2 failed: %v", err2)
	} else if info2.TF != 3 || info2.DF != 3 {
		t.Errorf("Test case 2 failed: expected TF=3, DF=3, got TF=%d, DF=%d", info2.TF, info2.DF)
	}

	info3, err3 := SearchWord("orange", fileContents)
	if err3 != nil {
		t.Errorf("Test case 3 failed: %v", err3)
	} else if info3.TF != 3 || info3.DF != 2 {
		t.Errorf("Test case 3 failed: expected TF=3, DF=2, got TF=%d, DF=%d", info3.TF, info3.DF)
	}

	info4, err1 := SearchWord("apple", fileContents)
	if err1 != nil {
		t.Errorf("Test case 4 failed: %v", err1)
	} else if info4.TF != 5 || info4.DF != 3 {
		t.Errorf("Test case 4 failed: expected TF=5, DF=3, got TF=%d, DF=%d", info4.TF, info4.DF)
	} else if info4.SearchCount != 2 {
		t.Errorf("Test case 4 failed: expected SearchCount=2, got SearchCount=%d", info4.SearchCount)
	}

}

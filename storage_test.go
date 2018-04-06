package mcdf

import (
	"testing"
)

func TestSetEntryName(t *testing.T) {
	entry := new(DirectoryEntry)
	testChars := "/"
	err := entry.SetEntryName(testChars)
	if err == nil || err != ErrIllegalCharacters {
		t.Errorf("illegal character not detected %v", testChars)
	}

	testChars = "\\"
	err = entry.SetEntryName(testChars)
	if err == nil || err != ErrIllegalCharacters {
		t.Errorf("illegal character not detected %v", testChars)
	}

	testChars = ":"
	err = entry.SetEntryName(testChars)
	if err == nil || err != ErrIllegalCharacters {
		t.Errorf("illegal character not detected %v", testChars)
	}

	testChars = "!"
	err = entry.SetEntryName(testChars)
	if err == nil || err != ErrIllegalCharacters {
		t.Errorf("illegal character not detected %v", testChars)
	}

	testChars = "aaaaabbbbbcccccdddddeeeeefffffg"
	err = entry.SetEntryName(testChars)
	if err != nil {
		t.Errorf("31 character entry name detected as too long %v", testChars)
	}

	testChars = "aaaaabbbbbcccccdddddeeeeefffffgg"
	err = entry.SetEntryName(testChars)
	if err == nil || err != ErrEntryNameTooLong {
		t.Errorf("32 character entry name not detected as too long %v", testChars)
	}
}

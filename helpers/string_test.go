package helpers

import "testing"

func TestPrefix(t *testing.T) {
	expectedString := "prefixed-word"
	if stringHelper := AddPrefix("prefixed", "word"); stringHelper != expectedString {
		t.Errorf("Expected stringHelper to be %s but got %s", expectedString, stringHelper)
	}
}

func TestSufix(t *testing.T) {
	expectedString := "word-sufixed"
	if stringHelper := AddSufix("word", "sufixed"); stringHelper != expectedString {
		t.Errorf("Expected stringHelper to be %s but got %s", expectedString, stringHelper)
	}
}

package helpers

import (
	"testing"

	"github.com/omarvides/hellogo/helpers"
)

func TestPrefix(t *testing.T) {
	expectedString := "prefixed-word"
	if stringHelper := helpers.AddPrefix("prefixed", "word"); stringHelper != expectedString {
		t.Errorf("Expected stringHelper to be %s but got %s", expectedString, stringHelper)
	}
}

func TestGetSum(t *testing.T) {
	expectedValue := 8
	if sumFunc := helpers.GetSum(5); sumFunc(3) != expectedValue {
		t.Errorf("Expected sumFunc(3) when GetSum(5) to be %d but got %d", expectedValue, sumFunc(3))
	}
}

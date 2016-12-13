package helpers

import "testing"

func TestGetSum(t *testing.T) {
	expectedValue := 8
	if sumFunc := GetSum(5); sumFunc(3) != expectedValue {
		t.Errorf("Expected sumFunc(3) when GetSum(5) to be %d but got %d", expectedValue, sumFunc(3))
	}
}

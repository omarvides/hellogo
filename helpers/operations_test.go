package helpers

import "testing"

func TestGetSum(t *testing.T) {
	expectedValue := 8
	if sumFunc := GetSum(5); sumFunc(3) != expectedValue {
		t.Errorf("Expected sumFunc(3) when GetSum(5) to be %d but got %d", expectedValue, sumFunc(3))
	}
}

func TestGetSub(t *testing.T) {
	expectedValue := 2
	if subFunc := GetSub(5); subFunc(3) != expectedValue {
		t.Errorf("Expected sumFunc(3) when GetSum(5) to be %d but got %d", expectedValue, subFunc(3))
	}
}

func TestGetMul(t *testing.T) {
	expectedValue := 15
	if mulFunc := GetMul(5); mulFunc(3) != expectedValue {
		t.Errorf("Expected sumFunc(3) when GetSum(5) to be %d but got %d", expectedValue, mulFunc(3))
	}
}

func TestGetDiv(t *testing.T) {
	expectedValue := 0.6
	if divFunc := GetDiv(5); divFunc(3) != expectedValue {
		t.Errorf("Expected sumFunc(3) when GetSum(5) to be %d but got %d", expectedValue, divFunc(3))
	}
}

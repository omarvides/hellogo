package helpers

import (
	"testing"

	"github.com/omarvides/hellogo/helpers"
)

func TestMain(t *testing.T) {
	result := helpers.AddPrefix("-sufix")
	if result != "prefix-sufix" {
		t.Error("Expected 'prefix-sufix' and got", result)
	}
}

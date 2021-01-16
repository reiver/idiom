package idiom_string

import (
	"testing"
)

func TestNothing(t *testing.T) {

	var uninitialized Type
	nothing := Nothing()

	if nothing != uninitialized {
		t.Error("An uninitialized idiom_string.Type is not equal to idiom_string.Nothing()")
		t.Logf("       uninitialized -> %#v", uninitialized)
		t.Logf("idiom_string.Nothing -> %#v", nothing)
		return
	}
}

package idiom_kernel

import (
	"testing"
)

func TestInternalRegistrar(t *testing.T) {

	var registry Registrar = &internalRegistrar{} // THIS LINE IS WHAT ACTUALLY MATTERS!

	if nil == registry {
		t.Errorf("This should never happen!")
	}
}


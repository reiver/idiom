package idiom_kernel

import (
	"testing"
)

func TestHandlerNotFound(t *testing.T) {

	var err error = errHandlerNotFound("does-not-matter") // THIS LINE IS WHAT ACTUALLY MATTERS!

	if nil == err {
		t.Errorf("This should never happen!")
	}
}

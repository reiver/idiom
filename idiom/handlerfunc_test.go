package idiom

import (
	"testing"
)

func TestHandlerFuncHandler(t *testing.T) {

	var handler Handler = HandlerFunc(nil) // THIS LINE IS WHAT ACTUALLY MATTERS!

	if nil == handler {
		t.Errorf("This should never happen!")
	}
}

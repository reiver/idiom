package idiom

import (
	"errors"
)

var (
	errHandlerFound = errors.New("idiom: Handler Found")
	errNilHandler   = errors.New("idiom: Nil Handler")
	errNilReceiver  = errors.New("idiom: Nil Receiver")
)

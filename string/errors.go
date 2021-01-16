package idiom_string

import (
	"errors"
)

var (
	errInternalError = errors.New("idiom: internal error")
	ErrNothing = errors.New("idiom_string.Nothing()")
)

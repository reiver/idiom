package idiom_token

import (
	"errors"
)

var (
	ErrEOF       = errors.New("end of file")
	ErrEOL       = errors.New("end of line")
	ErrUndefined = errors.New("undefined")
)

var (
	errInternalError = errors.New("internal error")
)

package idiom_tokenizer

import (
	"errors"
)

var (
	errNilReceiver    = errors.New("nil receiver")
	errNilRuneScanner = errors.New("nil rune scanner")
)

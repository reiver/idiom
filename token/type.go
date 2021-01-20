/*
Package idiom_token provides a type that represents an idiom token.

*/
package idiom_token

import (
	"fmt"
)

const (
	undefined = 0
	str       = 1
	eol       = 2
	eof       = 3
)

type Type struct {
	value string
	kind  int
}

func EOF() Type {
	return Type{
		kind:eof,
	}
}

func EOL() Type {
	return Type{
		kind:eol,
	}
}

func Undefined() Type {
	return Type{
		kind:undefined,
	}
}

func String(value string) Type {
	return Type{
		kind:str,
		value:value,
	}
}

func (receiver Type) Value() string {
	return receiver.value
}

func (receiver Type) Return() (string, error) {
	switch {
	case Undefined() == receiver:
                return "", ErrUndefined
	case EOL() == receiver:
                return "", ErrEOL
	case EOF() == receiver:
                return "", ErrEOF
	case str == receiver.kind:
		return receiver.value, nil
        default:
                return "", errInternalError
	}
}

func (receiver Type) GoString() string {
	switch {
	case Undefined() == receiver:
		return "idiom_token.Undefined()"
	case EOL() == receiver:
		return "idiom_token.EOL()"
	case EOF() == receiver:
		return "idiom_token.EOF()"
	case str == receiver.kind:
		return fmt.Sprintf("idiom_token.String(%q)", receiver.value)
	default:
		return fmt.Sprintf("#%v", receiver)
	}
}

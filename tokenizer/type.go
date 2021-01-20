package idiom_tokenizer

import (
	"../token"

	"bufio"
	"io"
)

type Type struct {
	tokens []idiom_token.Type
	scanner io.RuneScanner
}

func New(reader io.Reader) *Type {
	return &Type{
		scanner: bufio.NewReader(reader),
	}
}

func (receiver *Type) PopToken() (idiom_token.Type, error) {
	if nil == receiver {
		return idiom_token.Undefined(), errNilReceiver
	}

	var value idiom_token.Type

	switch {
	case 0 < len(receiver.tokens):
		value, receiver.tokens = receiver.tokens[len(receiver.tokens)-1], receiver.tokens[:len(receiver.tokens)-1]
	default:
		var err error
		value, err = readtoken(receiver.scanner)
		if nil != err {
			return idiom_token.Undefined(), err
		}
	}

	return value, nil
}

func (receiver *Type) PushToken(value idiom_token.Type) {
	if nil == receiver {
		panic("nil receiver")
	}

	receiver.tokens = append(receiver.tokens, value)
}

package idiom_parser

import (
	"../token"
	"../tokenizer"

	"io"
)

type Type struct {
	tokenizer *idiom_tokenizer.Type
}

func New(reader io.Reader) *Type {
	return &Type{
		tokenizer: idiom_tokenizer.New(reader),
	}
}

func (receiver *Type) Readln() ([]idiom_token.Type, error) {
	if nil == receiver {
		return nil, errNilReceiver
	}

	var tokens []idiom_token.Type

	Loop:
	for {
		token, err := receiver.tokenizer.PopToken()
		if nil != err {
			return nil, err
		}
		switch {

		case idiom_token.EOF() == token:
			if 0 >= len(tokens) {
				return []idiom_token.Type{token}, nil
			}
	/////////////// BREAK
			break Loop

		case idiom_token.EOL() == token:
			if 0 >= len(tokens) {
				return []idiom_token.Type(nil), nil
			}
	/////////////// BREAK
			break Loop

		case idiom_token.String("|") == token:
			if 0 < len(tokens) {
				receiver.tokenizer.PushToken(token)
	/////////////////////// BREAK
				break Loop
			}
			tokens = append(tokens, token)

		default:
			tokens = append(tokens, token)
		}
	}

	return tokens, nil
}

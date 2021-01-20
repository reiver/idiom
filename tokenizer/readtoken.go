package idiom_tokenizer

import (
	"../token"

	"errors"
	"io"
	"strings"
	"unicode"
)

func readtoken(scanner io.RuneScanner) (idiom_token.Type, error) {
	if nil == scanner {
		return idiom_token.Undefined(), errNilRuneScanner
	}

	if err := discardWhitespace(scanner); nil != err {
		return idiom_token.Undefined(), err
	}

	var buffer strings.Builder

	Loop:
	for {
		r, _, err := scanner.ReadRune()
		if io.EOF == err {
			if 0 == buffer.Len() {
				return idiom_token.EOF(), nil
			}
	/////////////// BREAK
			break Loop
		}
		if nil != err {
			return idiom_token.Undefined(), err
		}
		switch r {
		case
			'\u0009', // horizontal tab
			'\u0020', // space
			'\u00A0', // no-break space
			'\u1680', // ogham space mark
			'\u180E', // mongolian vowel separator
			'\u2000', // en quad
			'\u2001', // em quad
			'\u2002', // en space
			'\u2003', // em space
			'\u2004', // three-per-em space
			'\u2005', // four-per-em space
			'\u2006', // six-per-em space
			'\u2007', // figure space
			'\u2008', // punctuation space
			'\u2009', // thin space
			'\u200A', // hair space
			'\u202F', // narrow no-break space
			'\u205F', // medium mathematical space
			'\u3000': // ideographic space
	/////////////// BREAK
			break Loop
		case
			'\u000A', // line feed
			'\u000B', // vertical tab
			'\u000C', // form feed
			'\u000D', // carriage return
			'\u0085', // next line
			'\u2028', // line separator
			'\u2029': // paragraph separator
			if 0 == buffer.Len() {
				return idiom_token.EOL(), nil
			}
			scanner.UnreadRune()
	/////////////// BREAK
			break Loop
		}

		buffer.WriteRune(r)
	}

	return idiom_token.String(buffer.String()), nil
}


func discardWhitespace(scanner io.RuneScanner) error {

	if nil == scanner {
		panic("idiom: internal error: nil scanner in discardWhitespace()")
	}

	for {
		r, _, err := scanner.ReadRune()

		if io.EOF == err {
			return nil
		}
		if nil != err {
			return err
		}
		if unicode.ReplacementChar == r {
			return errors.New("rune error")
		}

		switch r {
		case
			'\u0009', // horizontal tab
			'\u0020', // space
			'\u00A0', // no-break space
			'\u1680', // ogham space mark
			'\u180E', // mongolian vowel separator
			'\u2000', // en quad
			'\u2001', // em quad
			'\u2002', // en space
			'\u2003', // em space
			'\u2004', // three-per-em space
			'\u2005', // four-per-em space
			'\u2006', // six-per-em space
			'\u2007', // figure space
			'\u2008', // punctuation space
			'\u2009', // thin space
			'\u200A', // hair space
			'\u202F', // narrow no-break space
			'\u205F', // medium mathematical space
			'\u3000': // ideographic space
	/////////////// CONTINUE
			continue
		case
			'\u000A', // line feed
			'\u000B', // vertical tab
			'\u000C', // form feed
			'\u000D', // carriage return
			'\u0085', // next line
			'\u2028', // line separator
			'\u2029': // paragraph separator
			return scanner.UnreadRune()
		default:
			return scanner.UnreadRune()
		}
	}
}

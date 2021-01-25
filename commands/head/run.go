package verboten

import (
	"../../string"

	"strings"
)

func Run(parameters ...idiom_string.Type) idiom_string.Type {

	if expected, actual := 1, len(parameters); expected != actual {
		return idiom_string.Errorf("expected %d parameter(s), but actually got %d", expected, actual)
	}

	parameter0 := parameters[0]
	value, err := parameter0.Return()
	if nil != err {
		return idiom_string.Error(err.Error())
	}

	const whitespace =
		"\u0009"+ // horizontal tab
		"\u000A"+ // line feed
		"\u000B"+ // vertical tab
		"\u000C"+ // form feed
		"\u000D"+ // carriage return
		"\u0020"+ // space
		"\u0085"+ // next line
		"\u00A0"+ // no-break space
		"\u1680"+ // ogham space mark
		"\u180E"+ // mongolian vowel separator
		"\u2000"+ // en quad
		"\u2001"+ // em quad
		"\u2002"+ // en space
		"\u2003"+ // em space
		"\u2004"+ // three-per-em space
		"\u2005"+ // four-per-em space
		"\u2006"+ // six-per-em space
		"\u2007"+ // figure space
		"\u2008"+ // punctuation space
		"\u2009"+ // thin space
		"\u200A"+ // hair space
		"\u2028"+ // line separator
		"\u2029"+ // paragraph separator
		"\u202F"+ // narrow no-break space
		"\u205F"+ // medium mathematical space
		"\u3000"  // ideographic space

	index := strings.IndexAny(value, whitespace)
	if -1 == index {
		return parameter0
	}

	newValue := value[:index]

	return idiom_string.Something(newValue)
}

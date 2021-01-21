package verboten

import (
	"../../string"

	"strings"
)

func Run(parameters ...idiom_string.Type) idiom_string.Type {

	var builder strings.Builder

	if maxExpected, actual := 1, len(parameters); maxExpected < actual {
		return idiom_string.Errorf("expected at most %d parameter(s), but actually got %d", maxExpected, actual)
	} else if 1 == actual {
		parameter0 := parameters[0]

		switch {
		case parameter0.IsError():
			return parameter0
		case parameter0.IsNothing():
			// Nothing here.
		default:
			value, err := parameter0.Return()
			if nil != err {
				return idiom_string.Error(err.Error()) //@TODO: should we return the raw error?
			}
			builder.WriteString(value)
		}
	}

	builder.WriteRune(0x0A) // ASCII Line Feed

	return idiom_string.Something(builder.String())
}

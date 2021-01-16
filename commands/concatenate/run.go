package verboten

import (
	"../../string"

	"strings"
)

func Run(parameters ...idiom_string.Type) idiom_string.Type {

	var builder strings.Builder

	for _, s := range parameters {
		if idiom_string.Nothing() == s {
			continue
		}

		str, err := s.Unwrap()
		if nil != err {
			//@TODO: should we return all the errors?
			return s
		}

		builder.WriteString(str)
	}

	return idiom_string.Something(builder.String())
}

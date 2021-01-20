package verboten

import (
	"../../string"

	"strings"
)

func Run(parameters ...idiom_string.Type) idiom_string.Type {

	var builder strings.Builder

	for _, s := range parameters {
		switch {
		case s.IsNothing():
			continue
		case s.IsError():
			//@TODO: should we return all the errors?
			return s
		case s.IsSomething():
			if 0 < builder.Len() {
				builder.WriteRune(' ')
			}
			value, _ := s.Return()
			builder.WriteString(value)
		default:
			return idiom_string.Error("internal error")
		}
	}

	return idiom_string.Something(builder.String())
}

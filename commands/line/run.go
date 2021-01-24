package verboten

import (
	"../../string"

	"strings"
	"unicode/utf8"
)

func Run(parameters ...idiom_string.Type) idiom_string.Type {

	var builder strings.Builder
	var wasMandatoryBreak bool

	for _, s := range parameters {
		switch {
		case s.IsNothing():
			continue
		case s.IsError():
			//@TODO: should we return all the errors?
			return s
		case s.IsSomething():
			if 0 < builder.Len() && !wasMandatoryBreak {
				builder.WriteRune(' ')
			}
			value, _ := s.Return()
			builder.WriteString(value)
			func(){
				r, _ := utf8.DecodeLastRuneInString(value)
				if r == utf8.RuneError {
					return
				}
				switch r {
				case
					'\u000A', // line feed
					'\u000B', // vertical tab
					'\u000C', // form feed
					'\u000D', // carriage return
					'\u0085', // next line
					'\u2028', // line separator
					'\u2029': // paragraph separator
					wasMandatoryBreak = true
				default:
					wasMandatoryBreak = false
				}
			}()
		default:
			return idiom_string.Error("internal error")
		}
	}

	return idiom_string.Something(builder.String())
}

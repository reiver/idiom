package verboten

import (
	"../../kernel"
	"../../string"
)

func Run(parameters ...idiom_string.Type) idiom_string.Type {

	if 2 > len(parameters) {
		return idiom_string.Error("too few parameters")
	}
	if 2 < len(parameters) {
		return idiom_string.Error("too many parameters")
	}

        var wrappedValue idiom_string.Type = parameters[0]
	if !wrappedValue.IsSomething() {
		return wrappedValue
	}
	value, _ := wrappedValue.Return()

        var wrappedName  idiom_string.Type = parameters[1]
	if !wrappedName.IsSomething() {
		return wrappedName
	}
	name, _ := wrappedName.Return()

	return idiom_kernel.Stem.Variables().Set(name, value)
}

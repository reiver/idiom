package verboten

import (
	"../../kernel"
	"../../string"
)

func Run(parameters ...idiom_string.Type) idiom_string.Type {

	if 1 > len(parameters) {
		return idiom_string.Error("too few parameters")
	}
	if 1 < len(parameters) {
		return idiom_string.Error("too many parameters")
	}

        var wrappedName  idiom_string.Type = parameters[0]
	if !wrappedName.IsSomething() {
		return wrappedName
	}
	name, _ := wrappedName.Return()

	return idiom_kernel.Stem.Variables().Get(name)
}

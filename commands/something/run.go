package verboten

import (
	"../../string"
)

func Run(parameters ...idiom_string.Type) idiom_string.Type {

	if expected, actual := 1, len(parameters); expected != actual {
		return idiom_string.Errorf("expected %d parameter(s), but actually got %d", expected, actual)
	}
	parameter0 := parameters[0]

	if !parameter0.IsSomething() {
		return idiom_string.Error("not something")
	}

	return parameter0
}


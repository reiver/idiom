package verboten

import (
	"../../kernel"
	"../../string"

	"fmt"
	"io"
)

func Run(parameters ...idiom_string.Type) idiom_string.Type {
	return run(idiom_kernel.Stem.Stdout(), parameters...)
}

func run(stdout io.Writer, parameters ...idiom_string.Type) idiom_string.Type {

	if expected, actual := 1, len(parameters); expected != actual {
		return idiom_string.Errorf("expected %d parameter(s), but actually got %d", expected, actual)
	}
	parameter0 := parameters[0]

	if parameter0.IsSomething() {
		value, err := parameter0.Return()
		if nil == err {
			fmt.Fprint(stdout, value)
		}
	}

	return parameter0
}

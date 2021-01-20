package verboten

import (
	"../../kernel"
	log "../../log"
	"../../string"
)

func Run(parameters ...idiom_string.Type) idiom_string.Type {

	if 0 >= len(parameters) {
		return idiom_string.Error("no command to pipe to")
	}

	parameter0 := parameters[0]
	parameters = append([]idiom_string.Type{idiom_kernel.Stem.Returned()}, parameters[1:]...)

	var command string
	{
		var err error

		command, err = parameter0.Return()
		if nil != err {
			return idiom_string.Error(err.Error())
		}
	}

	if log.Debugging() {
		log.Debug("[pipe] BEGIN")
		for i, p := range parameters {
			log.Debugf("parameters[%d] = %#v", i, p)
		}
		log.Debug("[pipe] END")
	}

	return idiom_kernel.Stem.Run(command, parameters...)
}

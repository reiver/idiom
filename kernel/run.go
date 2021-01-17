package idiom_kernel

import (
	"../string"
)

func Run(command string, parameters ...idiom_string.Type) idiom_string.Type {
	handler, err := Commands.Fetch(command)
	if nil != err {
		return idiom_string.Error(err.Error()) //@TODO: Should it send the raw error message?
	}

	return handler.Run(parameters...)
}

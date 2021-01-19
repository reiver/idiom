package idiom_kernel

import (
	"../string"
)

func (receiver *Branch) Run(command string, parameters ...idiom_string.Type) idiom_string.Type {
        if nil == receiver {
                return idiom_string.Error("nil receiver")
        }

	handler, err := receiver.Commands().Fetch(command)
	if nil != err {
		return idiom_string.Error(err.Error()) //@TODO: Should it send the raw error message?
	}

	return handler.Run(parameters...)
}

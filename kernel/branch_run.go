package idiom_kernel

import (
	log "../log"
	"../string"
)

func (receiver *Branch) Run(command string, parameters ...idiom_string.Type) idiom_string.Type {

	if log.Tracing() {
		log.Trace("[run] BEGIN")
		log.Tracef("[run] command = %q", command)
		log.Tracef("[run] PREVIOUS returned = %#v", receiver.Returned())
		log.Tracef("[run] len(parameters) = %d", len(parameters))
		for i,p := range parameters {
			log.Tracef("[run]\tparameters[%d] = %#v", i, p)
		}
	}

	if nil == receiver {
		log.Trace("[run] END ERROR (nil receiver)")
		return idiom_string.Error("nil receiver")
	}

	handler, err := receiver.Commands().Fetch(command)
	if nil != err {
		log.Tracef("[run] END ERROR (%s)", err.Error())
		return idiom_string.Error(err.Error()) //@TODO: Should it send the raw error message?
	}

	returned := handler.Run(parameters...)
	receiver.returned = returned
	log.Tracef("[run] returned = %#v", returned)

	log.Trace("[run] END")
	return returned
}

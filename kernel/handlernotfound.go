package idiom_kernel

import (
	"fmt"
)

type HandlerNotFound interface {
	error
	HandlerNotFound() string
}

type internalHandlerNotFound struct {
	name string
}

func errHandlerNotFound(name string) error {
	var e HandlerNotFound = &internalHandlerNotFound{
		name: name,
	}

	return e
}

func (receiver internalHandlerNotFound) Error() string {
	return fmt.Sprintf("idiom: handler not found for name=%q", receiver.name)
}

func (receiver internalHandlerNotFound) HandlerNotFound() string {
	return receiver.name
}

package idiom_kernel

import (
	"../string"
)

func (receiver Branch) Returned() idiom_string.Type {
	return receiver.returned
}

package idiom_kernel

import (
	"../string"
)

type Branch struct {
	returned  idiom_string.Type
	commands  internalRegistrar
	variables internalConstructor
}

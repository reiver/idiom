package idiom_kernel

import (
	"../string"
)

//  HandlerFunc is an adaptor that turns a func with signature:…
//
//	func(...idiom_string.Type)idiom_string.Type
//
// … into a idioms.Handler
type HandlerFunc func(...idiom_string.Type)idiom_string.Type

func (fn HandlerFunc) Run(parameters ...idiom_string.Type) idiom_string.Type {
	return fn(parameters...)
}

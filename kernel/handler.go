package idioms

import (
	"../string"
)

// Handler represents the implementation side — what actually does the work — of a specific idiom command.
//
// NOTE that the name of the idiom command is not part of the Handler. But instead the name of the command
// is set when the Handler is registered.
type Handler interface {
	Run(...idiom_string.Type) idiom_string.Type
}

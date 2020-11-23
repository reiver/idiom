package verboten

import (
	"../../idiom"
)

func init() {
	idiom.Commands.Register(idiom.HandlerFunc(Run), name)
}

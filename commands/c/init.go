package verboten

import (
	"../../idiom"
)

func init() {
	if err := idiom.Commands.Register(idiom.HandlerFunc(Run), name); nil != err {
		panic(err)
	}
}

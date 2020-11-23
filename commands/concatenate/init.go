package verboten

import (
	"../../idiom"
)

func init() {
	for _, name := range names {
		if err := idiom.Commands.Register(idiom.HandlerFunc(Run), name); nil != err {
			panic(err)
		}
	}
}

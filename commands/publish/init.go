package verboten

import (
	"../../kernel"
)

func init() {
	for _, name := range names {
		if err := idiom_kernel.Commands.Register(idiom_kernel.HandlerFunc(Run), name); nil != err {
			panic(err)
		}
	}
}

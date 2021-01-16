package verboten

import (
	"../../idioms"
)

func init() {
	for _, name := range names {
		if err := idioms.Commands.Register(idioms.HandlerFunc(Run), name); nil != err {
			panic(err)
		}
	}
}

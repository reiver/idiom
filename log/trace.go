package idiom_log

import (
	"../cmdln"

	"fmt"
)

func Tracing() bool {
	return idiom_cmdln.VeryVeryVerbose
}

func Trace(a ...interface{}) {
	if !Tracing() {
		return
	}

	s := fmt.Sprint(a...)

	Tracef("%s", s)
}

func Tracef(format string, a ...interface{}) {
	if !Tracing() {
		return
	}

	format = fmt.Sprintf("%s%s%s%s%s", "\x1b[48;2;1;1;1m", "\x1b[38;2;255;199;6m", format, "\x1b[0m", "\n")
	fmt.Printf(format, a...)
}

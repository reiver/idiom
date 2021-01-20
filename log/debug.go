package idiom_log

import (
	"../cmdln"

	"fmt"
)

func Debugging() bool {
	return idiom_cmdln.VeryVeryVerbose || idiom_cmdln.VeryVerbose
}

func Debug(a ...interface{}) {
	if !Debugging() {
		return
	}

	s := fmt.Sprint(a...)

	Debugf("%s", s)
}

func Debugf(format string, a ...interface{}) {
	if !Debugging() {
		return
	}

	format = fmt.Sprintf("%s%s%s%s%s", "\x1b[48;2;1;1;1m", "\x1b[38;2;44;181;233m", format, "\x1b[0m", "\n")
	fmt.Printf(format, a...)
}

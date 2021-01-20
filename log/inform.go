package idiom_log

import (
	"../cmdln"

	"fmt"
)

func Informing() bool {
	return idiom_cmdln.VeryVeryVerbose || idiom_cmdln.VeryVerbose || idiom_cmdln.Verbose
}

func Inform(a ...interface{}) {
	if !Informing() {
		return
	}

	s := fmt.Sprint(a...)

	Informf("%s", s)
}

func Informf(format string, a ...interface{}) {
	if !Informing() {
		return
	}

	format = fmt.Sprintf("%s%s%s%s%s%s", "\x1b[48;2;1;1;1m", "\x1b[38;222;56;43;6m", format, "\x1b[0m", "\x1b[0m", "\n")
	fmt.Printf(format, a...)
}

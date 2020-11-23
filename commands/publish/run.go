package verboten

import (
	"../../idiom"

	"fmt"
	"io"
	"strings"
)

func Run(parameters ...string) (string, error) {
	return run(idiom.Stdout, parameters...)
}

func run(stdout io.Writer, parameters ...string) (string, error) {

	var buffer strings.Builder

	buffer.WriteRune('«')

	for i,s := range parameters {
		if "" == s {
			continue
		}
		if 0 < i {
			buffer.WriteRune(' ')
		}
		buffer.WriteString(s)
	}

	buffer.WriteRune('»')

	fmt.Fprintln(stdout, buffer.String())
	return "", nil
}

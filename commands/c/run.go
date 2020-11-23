package verboten

import (
	"strings"
)

func Run(parameters ...string) (string, error) {

	var builder strings.Builder

	for _, s := range parameters {
		builder.WriteString(s)
	}

	return builder.String(), nil
}

package idiom

import (
	"fmt"
)

func Errorf(format string, a ...interface{}) error {
	const prefix = "idiom"

	format = fmt.Sprintf("%s: %s", prefix, format)

	return fmt.Errorf(format, a...)
}

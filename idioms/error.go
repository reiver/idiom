package idiom

import (
	"errors"
	"fmt"
)

const errorPrefix = "idiom"

func Error(a ...interface{}) error {
	prefix := fmt.Sprintf("%s:", errorPrefix)

	a = append([]interface{}{prefix}, a...)

	text := fmt.Sprint(a...)

	return errors.New(text)
}

func Errorf(format string, a ...interface{}) error {
	format = fmt.Sprintf("%s: %s", errorPrefix, format)

	return fmt.Errorf(format, a...)
}

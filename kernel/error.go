package idioms

import (
	"errors"
	"fmt"
)

const errorPrefix = "idiom"

// Error returns an error whose message is prefixed with "idiom: ".
//
// For example with:…
//
//	err := idioms.Error("something bad happened")
//
// … the errors message would be:
//
//	idiom: something bad happened
func Error(a ...interface{}) error {
	prefix := fmt.Sprintf("%s:", errorPrefix)

	a = append([]interface{}{prefix}, a...)

	text := fmt.Sprint(a...)

	return errors.New(text)
}

// Errorf returns an error whose message is prefixed with "idiom: ".
//
// For example with:…
//
//	var name = "Joe Blow"
//	
//	err := idioms.Errorf("%s is not allowed here", name)
//
// … the errors message would be:
//
//	idiom: Joe Blow is not allowed here
func Errorf(format string, a ...interface{}) error {
	format = fmt.Sprintf("%s: %s", errorPrefix, format)

	return fmt.Errorf(format, a...)
}

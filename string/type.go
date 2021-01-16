package idiom_string

import(
	"errors"
	"fmt"
)

const (
	stateNothing   = 0
	stateSomething = 1
	stateError     = 2
)

type Type struct {
	state  int
	value  string
}

func Nothing() Type {
	return Type{} // ‘state’ defaults to ‘stateNothing’ since stateNothing == 0
}

func Error(err string) Type {
	return Type{
		state: stateError,
		value:   err,
	}
}

func Something(value string) Type {
	return Type{
		state:  stateSomething,
		value:  value,
	}
}

func (receiver Type) IsNothing() bool {
	return stateNothing == receiver.state
}

func (receiver Type) IsError() bool {
	return stateError == receiver.state
}

func (receiver Type) IsSomething() bool {
	return stateSomething == receiver.state
}

func (receiver Type) GoString() string {
	switch {
	case receiver.IsNothing():
		return "idiom_string.Nothing()"
	case receiver.IsError():
		return fmt.Sprintf("idiom_string.Error(%q)", receiver.value)
	case receiver.IsSomething():
		return fmt.Sprintf("idiom_string.Something(%q)", receiver.value)
	default:
		return fmt.Sprintf("#%v", receiver)
	}
}

func (receiver Type) Return() (string, error) {
	switch {
	case receiver.IsNothing():
		return "", ErrNothing
	case receiver.IsError():
		return "", errors.New(receiver.value)
	case receiver.IsSomething():
		return receiver.value, nil
	default:
		return "", errInternalError
	}
}



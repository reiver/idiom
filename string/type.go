package idiom_string

import(
	"errors"
)

const (
	stateNothing   = 0
	stateSomething = 1
	stateError     = 2
)

type Type struct {
	state  int
	value  string
	err    string
}

func Nothing() Type {
	return Type{} // ‘state’ defaults to ‘stateNothing’ since stateNothing == 0
}

func Error(err string) Type {
	return Type{
		state: stateError,
		err:   err,
	}
}

func Something(value string) Type {
	return Type{
		state:  stateSomething,
		value:  value,
	}
}

func (receiver Type) Unwrap() (string, error) {
	if Nothing() == receiver {
		return "", ErrNothing
	}
	if stateError == receiver.state {
		return "", errors.New(receiver.err)
	}
	return receiver.value, nil
}

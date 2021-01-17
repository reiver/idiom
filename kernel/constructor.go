package idiom_kernel

import (
	"../string"

	"sync"
)

type Constructor interface {
	Get(string) idiom_string.Type
	Set(string, string) idiom_string.Type
	Unset(string) idiom_string.Type
}

type internalConstructor struct {
	mutex sync.RWMutex
	values map[string]string
}

func (receiver *internalConstructor) Get(name string) idiom_string.Type {
	if nil == receiver {
		panic(errNilReceiver)
	}

	receiver.mutex.RLock()
	defer receiver.mutex.RUnlock()

	values := receiver.values
        if nil == values {
		return idiom_string.Nothing()
        }

        value, found := values[name]
        if !found {
                return idiom_string.Nothing()
        }

        return idiom_string.Something(value)
}

func (receiver *internalConstructor) Set(name string, value string) idiom_string.Type {
	if nil == receiver {
		panic(errNilReceiver)
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	if nil == receiver.values {
		receiver.values = map[string]string{}
	}

	receiver.values[name] = value

	return idiom_string.Something(value)
}

func (receiver *internalConstructor) Unset(name string) idiom_string.Type {
	if nil == receiver {
		panic(errNilReceiver)
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	if nil == receiver.values {
		return idiom_string.Nothing()
	}

	value, found := receiver.values[name]
	delete(receiver.values, name)

	if !found {
		return idiom_string.Nothing()
	}
	return idiom_string.Something(value)
}

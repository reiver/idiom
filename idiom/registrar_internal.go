package idiom

import (
	"sync"
)


type internalRegistrar struct {
	mutex sync.RWMutex
	handlers  map[string]Handler
	summaries map[string]string
}

func (receiver *internalRegistrar) Fetch(name string) (Handler, error) {
	if nil == receiver {
		return nil, errNilReceiver
	}

	receiver.mutex.RLock()
	defer receiver.mutex.RUnlock()

	handlers := receiver.handlers
	if nil == handlers {
		return nil, errHandlerNotFound(name)
	}

	handler, found := handlers[name]
	if !found {
		return nil, errHandlerNotFound(name)
	}
	if nil == handler {
		return nil, errNilHandler
	}

	return handler, nil
}

func (receiver *internalRegistrar) Register(handler Handler, name string) error {
	if nil == receiver {
		return errNilReceiver
	}

	if nil == handler {
		return errNilHandler
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	if nil == receiver.handlers {
		receiver.handlers = map[string]Handler{}
	}

	_, found := receiver.handlers[name]
	if found {
		return errHandlerFound
	}

	receiver.handlers[name]  = handler

	return nil
}

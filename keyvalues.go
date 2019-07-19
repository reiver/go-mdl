package mdl

import (
	"sync"
)

// KeyValues represents key-value pairs.
//
// It stores key-value pairs.
//
// ‘mdl.KeyValues’ is the type of one of the fields for ‘mdl.Instruction’.
type KeyValues struct {
	mutex sync.RWMutex
	data map[Key]string
}

// Fetch is similar to .Load().
func (receiver *KeyValues) Fetch(key ...string) String {
	return receiver.Load(SomeKey(key...))
}

func (receiver *KeyValues) For(fn func(Key, string)) {
	if nil == receiver {
		return
	}

	receiver.mutex.RLock()
	defer receiver.mutex.RUnlock()

	for k, v := range receiver.data {
		fn(k, v)
	}

}

func (receiver *KeyValues) Len() int {
	if nil == receiver {
		return 0
	}

	receiver.mutex.RLock()
	defer receiver.mutex.RUnlock()

	return len(receiver.data)
}

func (receiver *KeyValues) Load(key Key) String {
	if nil == receiver {
		return NoString()
	}
	if NoKey() == key {
		return NoString()
	}

	receiver.mutex.RLock()
	defer receiver.mutex.RUnlock()

	data := receiver.data
	if nil == data {
		return NoString()
	}

	value, found := data[key]
	if !found {
		return NoString()
	}

	return SomeString(value)
}

// ShallowStore being called as:
//
//	err := keyvalues.ShallowStore(key, value)
//
// ... is equivalent to calling Store as:
//
//	err := keyvalues.Store(mdl.SomeKey(key), value)
func (receiver *KeyValues) ShallowStore(key string, value string) error {
	return receiver.Store(SomeKey(key), value)
}

func (receiver *KeyValues) Store(key Key, value string) error {
	if nil == receiver {
		return errNilReceiver
	}

	if NoKey() == key {
		return errEmptyKey
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	if nil == receiver.data {
		receiver.data = map[Key]string{}
	}

	foundValue, found := receiver.data[key]
	if found {
		return internalKeyFound{
			key:key,
			value:value,
			foundValue:foundValue,
		}
	}

	receiver.data[key] = value

	return nil
}

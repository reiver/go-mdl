package mdl

import (
	"fmt"
)

// KeyFound is the error returned from mdl.KeyValues.Store() when .Store() is called and there is
// already a value stored for that key.
//
// For example:
//
//	var keyvalues mdl.KeyValues
//	
//	// ...
//	
//	err := keyvalues.Store(value, key...)
//	
//	if nil != err {
//		switch err.(type) {
//		case mdl.KeyFound:
//			//@TODO
//		default:
//			//@TODO
//		}
//	}
type KeyFound interface {
	error
	KeyFound()

	// Key returns the key.
	Key() []string

	// Value returns the value that was trying to be stored at the key, when the error occurred.
	Value() string

	// FoundValue returns the value that was already stored at the key.
	FoundValue() string
}

type internalKeyFound struct {
	key []string
	value string
	foundValue string
}

func (receiver internalKeyFound) Error() string {
	return fmt.Sprintf("mdl: key %q found with value %q, cannot store value %q", receiver.key, receiver.foundValue, receiver.value)
}

func (receiver internalKeyFound) Key() []string {
	return receiver.key
}

func (receiver internalKeyFound) Value() string {
	return receiver.value
}

func (receiver internalKeyFound) FoundValue() string {
	return receiver.foundValue
}

func (internalKeyFound) KeyFound() {
	// Nothing here.
}

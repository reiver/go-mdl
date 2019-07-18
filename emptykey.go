package mdl

// EmptyKey is the error returned from mdl.KeyValues.Store() when .Store() is called with a key that
// has the value of ‘mdl.NoKey()’.
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
//		case mdl.EmptyKey:
//			//@TODO
//		default:
//			//@TODO
//		}
//	}
type EmptyKey interface {
	error
	EmptyKey()
}

type internalEmptyKey struct{}

func (receiver internalEmptyKey) Error() string {
	return "mdl: empty key"
}

func (internalEmptyKey) EmptyKey() {
	// Nothing here.
}

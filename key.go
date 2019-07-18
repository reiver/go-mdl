package mdl

// Key represents the ‘key’ in the key-value pairs.
//
// Because the ‘key’ isn't just a ‘string’, and is instead (conceptually) a ‘[]string’,
// ‘mdl.Key’ exists to abstract that complexity, and make the API for this package a bit
// easier to use.
//
// Example
//
// Here is an example of setting the value of ‘mdl.Key’ using a non-serialized format:
//
//	var key mdl.Key
//	
//	// ...
//	
//	key = mdl.SomKey("database", "password")
//
// Example
//
// Here is an example of setting the value of ‘mdl.Key’ using a serialized format:
//
//	var key mdl.Key
//	
//	// ...
//	
//	err := key.Scan("database/password")
type Key struct {
	encoded String
}

// NoKey returns a ‘mdl.Key’ which has no value (i.e., ‘nothing’).
//
// Example
//
// Here is an example of ‘mdl.NoKey()’ being used in an assignment.
//
//	var key mdl.Key
//	
//	// ...
//	
//	key = mdl.NoKey()
//
// You can also use ‘mdl.NoKey()’ in comparisons in an if-statment, as in for example:
//
//	var key mdl.Key
//	
//	// ...
//	
//	if mdl.NoKey() == key {
//		//@TODO
//	}
//
// And you can use ‘mdl.NoKey()’ in comparisons in an switch-statment, as in for example:
//
//	var key mdl.Key
//	
//	// ...
//	
//	switch key {
//	case mdl.NoKey():
//		//@TODO
//	
//	case mdl.SomeKey("database", "username"):
//		//@TODO
//	case mdl.SomeKey("database", "password"):
//		//@TODO
//	case mdl.SomeKey("version"):
//		//@TODO
//	
//	default:
//		//@TODO
//	}
func NoKey() Key {
	return Key{}
}

// SomeKey returns a ‘mdl.Key’ which has some value (i.e., ‘something’).
//
// Example
//
// Here is an example of ‘mdl.SomeKey()’ being used in an assignment.
//
//	var key mdl.Key
//	
//	// ...
//	
//	key = mdl.SomeKey("database", "password")
//
// You can also use ‘mdl.SomeKey()’ in comparisons in an if-statment, as in for example:
//
//	var key mdl.Key
//	
//	// ...
//	
//	if mdl.SomeKey("database", "password") == key {
//		//@TODO
//	}
//
// And you can use ‘mdl.SomeKey()’ in comparisons in an switch-statment, as in for example:
//
//	var key mdl.Key
//	
//	// ...
//	
//	switch key {
//	case mdl.NoKey():
//		//@TODO
//	
//	case mdl.SomeKey("database", "username"):
//		//@TODO
//	case mdl.SomeKey("database", "password"):
//		//@TODO
//	case mdl.SomeKey("version"):
//		//@TODO
//	
//	default:
//		//@TODO
//	}
func SomeKey(key ...string) Key {
	if 0 == len(key) {
		return NoKey()
	}

	encoded := KeySerialize(key...)

	return Key{
		encoded: SomeString(encoded),
	}
}

func (receiver Key) Else(key ...string) Key {
	if NoKey() != receiver {
		return receiver
	}

	return SomeKey(key...)
}

func (receiver Key) ElseUnwrap(key ...string) []string {
	if NoKey() == receiver {
		return key
	}

	a, _ := KeyDeserialize(receiver.encoded.datum)

	return a
}

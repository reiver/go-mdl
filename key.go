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

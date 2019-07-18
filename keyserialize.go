package mdl

import (
	"strings"
)

// KeySerialize returns the serliaized version of a ‘key’.
//
// In packge ‘mdl’, a ‘key’ isn't just a ‘string’, and is instead (conceptually) a ‘[]string’.
//
// Because many things cannot work with a ‘[]string’, we have a serliaized form of the ‘key’.
//
// Example
//
// Here is an example of using ‘mdl.KeySerialze()’.
//
//	serialized := mdl.KeySerialize("database", "password") // == "database/password"
//
// This is a more low-level function. Usually you will probably want to use the type ‘mdl.Key’.
func KeySerialize(key ...string) string {

	var builder strings.Builder

	for i, token := range key {
		if 0 != i {
			builder.WriteRune('/')
		}

		if err := keyTokenEncode(&builder, token); nil != err {
			return ""
		}
	}

	return builder.String()
}

package mdl

import (
	"strings"
	"unicode/utf8"
)

// KeyDeserialize returns the deserliaized version of a ‘serialized’.
//
// In packge ‘mdl’, a ‘key’ isn't just a ‘string’, and is instead (conceptually) a ‘[]string’.
//
// Because many things cannot work with a ‘[]string’, we have a serliaized form of the ‘key’.
//
// Example
//
// Here is an example of using ‘mdl.KeyDeserialze()’.
//
//	key := mdl.KeyDeserialize("database/password") // == []string{"database", "password"}
//
// This is a more low-level function. Usually you will probably want to use the type ‘mdl.Key’.
func KeyDeserialize(serialized string) ([]string, error) {
	if "" == serialized {
		return []string(nil), nil
	}

	var result []string

	var builder strings.Builder

	var escaped bool

	var s string = serialized
	for 0 < len(s) {
		r, size := utf8.DecodeRuneInString(s)
		if utf8.RuneError == r {
			return result, errRuneError
		}
		if 0 >= size {
			return result, errInternalError
		}
		s = s[size:]

		switch {
		case !escaped && '\\' == r:
			escaped = true
		case escaped:
			escaped = false
			builder.WriteRune(r)
		case '/' == r:
			result = append(result, builder.String())
			builder.Reset()
		default:
			builder.WriteRune(r)
		}

		if 0 == len(s) {
			result = append(result, builder.String())
		}
	}

	return result, nil
}

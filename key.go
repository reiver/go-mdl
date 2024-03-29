package mdl

import (
	"database/sql/driver"
	"fmt"
	"strings"
)

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

// CanonicalForm returns the ‘mdl.Key’ in ‘canonical form’.
//
// Example
//
// For example, for this ‘mdl.Key’:
//
//	var key mdl.Key = mdl.SomeKey("database", "password")
//
// ... its canonical form’ form is:
//
//	database/password
func (receiver Key) CanonicalForm() string {
	return receiver.encoded.ElseUnwrap("")
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

func (receiver Key) Format(f fmt.State, c rune) {
	switch c {
        case 'q':
		switch receiver {
		case NoKey():
			fmt.Fprint(f, "«no-key»")
		default:
			fmt.Fprintf(f, "%q", receiver.encoded)
		}
	default:
		fmt.Fprintf(f, "%%!%s(%s)", string(c), receiver.GoString())
	}
}

// GoString makes ‘mdl.Key’ fit the fmt.GoStringer interface.
//
// It gets used with the %#v verb with the printing family of functions
// in the Go built-in "fmt" package.
//
// I.e., it gets used with: fmt.Fprint(), fmt.Fprintf(), fmt.Fprintln(),
// fmt.Print(), fmt.Printf(), fmt.Println(), fmt.Sprint(), fmt.Sprintf(),
// fmt.Sprintln().
//
// Example
//
// Here is an example where .GoString() is being implicitly used.
// This implicit usage is the way most people are likely to use it.
//
//	var datum mdl.Key
//	
//	// ...
//	
//	fmt.Printf("datum = %#v\n", s) // <---- datum.GoString() is called by fmt.Printf()
func (receiver Key) GoString() string {
	if NoKey() == receiver {
		return "mdl.NoKey()"
	}

	var builder strings.Builder

	builder.WriteString("mdl.SomeKey(")
	for i, token := range receiver.ElseUnwrap() {
		if 0 != i {
			builder.WriteString(", ")
		}
		fmt.Fprintf(&builder, "%q", token)
	}
	builder.WriteRune(')')

	return builder.String()
}

func (receiver Key) Map(fn func(...string)[]string) Key {
	if NoKey() == receiver {
		return receiver
	}

//@TODO: Shouldn't this be an error?
	if nil == fn {
		return receiver
	}

	return SomeKey(fn(receiver.ElseUnwrap()...)...)
}

func (receiver Key) String() (string, error) {
	return receiver.encoded.String()
}

func (receiver Key) Then(fn func(...string)Key) Key {
	if NoKey() == receiver {
		return receiver
	}

	return fn(receiver.ElseUnwrap()...)
}

// Value makes mdl.Key fit the database/sql/driver.Valuer interface.
func (receiver Key) Value() (driver.Value, error) {
	if NoKey() == receiver {
		return nil, errNotLoaded
	}

	return receiver.encoded.datum, nil
}

package mdl

import (
	"database/sql/driver"
	"fmt"
)

// String is a string option type.
//
// It can have 2 kinds of value:
//
// • no string (i.e., ‘nothing’),
//
// • some string (i.e., ‘something’).
//
// The value in this type is that it can differentiate between being assigned an empty string (""),
// and a variable that hasn't been loaded.
type String struct {
	datum  string
	loaded bool
}

// NoString returns a mdl.String which has no value (i.e., ‘nothing’).
//
// Note that this is not the same thing as an empty string ("")!
// An empty string is mdl.SomeString("")
//
// Example
//
// Here is an example of mdl.NoString() being used in an assignment.
//
//	var datum mdl.String
//	
//	// ...
//	
//	datum = mdl.NoString()
//
// You can also use mdl.NoString() in comparisons in an if-statment, as in for example:
//
//	var datum mdl.String
//	
//	// ...
//	
//	if mdl.NoString == datum {
//		//@TODO
//	}
//
// And you can use mdl.NoString() in comparisons in an switch-statment, as in for example:
//
//	var datum mdl.String
//	
//	// ...
//	
//	switch datum {
//	case mdl.NoString():
//		//@TODO
//	
//	case mdl.SomeString("apple"):
//		//@TODO
//	case mdl.SomeString("banana"):
//		//@TODO
//	case mdl.SomeString("cherry"):
//		//@TODO
//	
//	default:
//		//@TODO
//	}
func NoString() String {
	return String{}
}

// SomeString returns a mdl.String which has some value (i.e., ‘something’).
//
// Example
//
// Here is an example of mdl.SomeString() being used in an assignment.
//
//	var datum mdl.String
//	
//	// ...
//	
//	datum = mdl.SomeString("Hello world!")
//
// You can also use mdl.SomeString() in comparisons in an if-statment, as in for example:
//
//	var datum mdl.String
//	
//	// ...
//	
//	if mdl.SomeString("Hello world!") == datum {
//		//@TODO
//	}
//
// And you can use mdl.SomeString() in comparisons in an switch-statment, as in for example:
//
//	var datum mdl.String
//	
//	// ...
//	
//	switch datum {
//	case mdl.NoString():
//		//@TODO
//	
//	case mdl.SomeString("apple"):
//		//@TODO
//	case mdl.SomeString("banana"):
//		//@TODO
//	case mdl.SomeString("cherry"):
//		//@TODO
//	
//	default:
//		//@TODO
//	}
func SomeString(datum string) String {
	return String{
		loaded: true,
		datum: datum,
	}
}

func (receiver String) Else(datum string) String {
	if NoString() != receiver {
		return receiver
	}

	return SomeString(datum)
}

func (receiver String) ElseUnwrap(datum string) string {
	if NoString() == receiver {
		return datum
	}

	return receiver.datum
}

// GoString makes mdl.String fit the fmt.GoStringer interface.
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
//	var datum mdl.String
//	
//	// ...
//	
//	fmt.Printf("datum = %#v\n", s) // <---- datum.GoString() is called by fmt.Printf()
func (receiver String) GoString() string {
	if NoString() == receiver {
		return "mdl.NoString()"
	}

	return fmt.Sprintf("mdl.SomeString(%q)", receiver.datum)
}

func (receiver String) Map(fn func(string)string) String {
	if NoString() == receiver {
		return receiver
	}

	return String {
		loaded: true,
		datum:  fn(receiver.datum),
	}
}

// Scan makes mdl.String fit the database/sql.Scanner interface.
func (receiver *String) Scan(src interface{}) error {
	if nil == receiver {
		return errNilReceiver
	}

	switch casted := src.(type) {
	case string:
		*receiver = SomeString(casted)
		return nil
	case []byte:
		switch casted {
		case nil:
			*receiver = NoString()
		default:
			*receiver = SomeString(string(casted))
		}
		return nil
	case String:
		*receiver = casted
		return nil
	case fmt.Stringer:
		*receiver = SomeString(casted.String())
		return nil
	case interface{String()(string,error)}:
		datum, err := casted.String()
		if nil != err {
			return err
		}
		*receiver = SomeString(datum)
		return nil
	default:
		return unsupportedSource(src)
	}
}

func (receiver String) String() (string, error) {
	if NoString() == receiver {
		return "", errNotLoaded
	}

	return receiver.datum, nil
}

func (receiver String) Then(fn func(string)String) String {
	if NoString() == receiver {
		return receiver
	}

	return fn(receiver.datum)
}

// Scan makes mdl.String fit the encoding.TextUnmarshaler interface.
func (receiver *String) UnmarshalText(text []byte) error {
	if nil == receiver {
		return errNilReceiver
	}

	return receiver.Scan(text)
}

func (receiver String) Unwrap() (string, bool) {
	return receiver.datum, receiver.loaded
}

// Value makes mdl.String fit the database/sql/driver.Valuer interface.
func (receiver String) Value() (driver.Value, error) {
	if NoString() == receiver {
		return receiver, errNotLoaded
	}

	return receiver.datum, nil
}

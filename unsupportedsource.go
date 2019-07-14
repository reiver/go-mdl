package mdl

import (
	"fmt"
	"runtime"
	"strings"
)

// UnsupportedSource is the error returned from mdl.String.Scan(), and mdl.String.UnmarshalText() when the type of the source is not supported.
//
// For example:
//
//	var source int64 // <---- Note that the type is int64, which mdl.String.Scan does not support (by itself). So it will return an error.
//	
//	// ...
//	
//	var datum mdl.String
//
//	// ...
//
//	err := datum.Scan(source)
//	
//	if nil != err {
//		switch err.(type) {
//		case mdl.UnsupportedSource:
//			//@TODO
//		default:
//			//@TODO
//		}
//	}
type UnsupportedSource interface {
	error

	// This UnsupportedSource() method exists to allow type checking of the error.
	UnsupportedSource()

	// Debug returns something similar to Error(), but also include more information useful for debugging, such as the file name, and line number the error came from.
	Debug() string

	// Source returns value that caused this error.
	Source() interface{}
}

type internalUnsupportedSource struct {
	source interface{}
	file   string
	line   int
}

func unsupportedSource(source interface{}) UnsupportedSource {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		line = -1
	}

	return internalUnsupportedSource{
		source: source,
		file:   file,
		line:   line,
	}
}

func (receiver internalUnsupportedSource) Debug() (s string) {
	defer func() {
		if r := recover(); nil != r {
			s = receiver.Error()
		}
	}()

	var file string = receiver.file
	{
		var index int = strings.LastIndex(file, "/")
		if 0 < index {
			file = file[1+index:]
		}
	}

	return fmt.Sprintf("mdl: %s:%d: unsupported source: %T", file, receiver.line, receiver.source)
}

func (receiver internalUnsupportedSource) Error() string {
	return fmt.Sprintf("mdl: unsupported source: %T", receiver.source)
}

func (receiver internalUnsupportedSource) Source() interface{} {
	return receiver.source
}

func (internalUnsupportedSource) UnsupportedSource() {
	// Nothing here.
}

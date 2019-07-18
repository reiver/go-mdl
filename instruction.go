package mdl

import (
	"net/http"
)

// Instruction represents a command, order, or direction.
//
// Some example instructions might be:
//
// • “empty that shopping cart”,
//
// • “add this book to that shopping cart”,
//
// • “add this item to that TODO list”, and
//
// • “add this e-mail address to my profile”.
//
// In CQRS terminology, a ‘mdl.Instruction’ would be the equivalent of a CQRS ‘command’ (i.e., the “C” in “CQRS”).
//
//
// HTTP Request
//
// Nowadays it is common to expose APIs over HTTP (or HTTPS).
//
// To help with this, ‘mdl.Instruction’ includes support for receiving ‘instructions’ over HTTP (or HTTPS).
//
// Since Go has the built-in "net/http" package for dealing with HTTP (and HTTPS), ‘mdl.Instruction’
// can infer the instruction from an ‘http.Request’.
//
// To do this, one would use mdl.Instruction's .Scan() method; doing something along the lines of:
//
//	func (receiver *MyHandler) ServeHTTP(responseWriter ResponseWriter, request *Request) {
//		
//		// ...
//		
//		var instruction mdl.Instruction
//	
//		// ...
//	
//		err := instruction.Scan(request)
//		switch err.(type) {
//		case mdl.BadRequest:
//			http.Error(responseWriter, "Bad Request", http.StatusBadRequest)
//			return
//		default:
//			http.Error(responseWriter, "Internal Server Error", http.StatusInternalServerError)
//			return
//		}
//		
//		// ...
//		
//	}
//
// This example Go source represents the ‘server’ side of the ‘instruction’.
// On the ‘client’ side of the ‘instruction’ we might have JavaScript code like the following:
//
//	httpRequest = new XMLHttpRequest();
//	
//	// ...
//	
//	httpRequest.open("POST", "https://api.example.com/v1/users");
//	
//	// This ‘X-Idempotent-ID’ HTTP request header would be different for each ‘instruction’.
//	//
//	// Although if, for example, you don't get a response back for an ‘instruction’, and
//	// you wanted to re-try the ‘instruction’, then you would use the same value for
//	// ‘X-Idempotent-ID’.
//	httpRequest.setRequestHeader("X-Idempotent-ID", "z-2015-05-07T10:25:09Z_tleEiguQe67zJFYUa7pngSZT8HX7FMAcHb1Z4yOO2ANtltRPRwF5p9TWwf7m");
//	
//	// Here we tell the server we are overriding the actual HTTP method used (i.e., "POST"
//	// in this example) to be instead be the non-common HTTP method “RECORD_EMAIL”.
//	httpRequest.setRequestHeader("X-HTTP-Method-Override", "RECORD_EMAIL");
//	
//	httpRequest.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
//	httpRequest.send('email_address=joeblow@example.com');
//
// Or this:
//
//	httpRequest = new XMLHttpRequest();
//	
//	// ...
//	
//	httpRequest.open("RECORD_EMAIL", "https://api.example.com/v1/users");
//	
//	// This ‘X-Idempotent-ID’ would be different for each ‘instruction’.
//	//
//	// Although if, for example, you don't get a response back for an ‘instruction’, and
//	// you wanted to re-try the ‘instruction’, then you would use the same value for
//	// ‘X-Idempotent-ID’.
//	httpRequest.setRequestHeader("X-Idempotent-ID", "z-2015-05-07T10:25:09Z_tleEiguQe67zJFYUa7pngSZT8HX7FMAcHb1Z4yOO2ANtltRPRwF5p9TWwf7m");
//	
//	httpRequest.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
//	httpRequest.send('email_address=joeblow@example.com');
//
//
// Idempotent ID
//
// IdempotentID is a unique identifier (ID) that helps make it so an instruction is executed at most one time.
//
// It helps make it so an instruction is idempotent.
//
//
// Idempotent ID From HTTP Request
//
// To provide an IdempotentID in an ‘http.Request’, have the HTTP client set the “X-Idempotent-ID” HTTP header.
//
// For example:
//
//	PATCH /v1/user/email HTTP/1.1
//	Host: api.example.com
//	X-Idempotent-ID: z-2015-05-07T10:25:09Z_tleEiguQe67zJFYUa7pngSZT8HX7FMAcHb1Z4yOO2ANtltRPRwF5p9TWwf7m
//	
//	email_address=joeblow@email.com
//
// If someone is using XMLHttpRequest, to make the HTTP request, then this can be set with code similar to:
//
//	httpRequest = new XMLHttpRequest();
//	
//	// ...
//	
//	httpRequest.httpRequest("X-Idempotent-ID", "z-2015-05-07T10:25:09Z_tleEiguQe67zJFYUa7pngSZT8HX7FMAcHb1Z4yOO2ANtltRPRwF5p9TWwf7m")
//
//
// How To Construct An Idempotent ID
//
// I recommend using a combination of the current time, and randomness be used to contruct Idempotent IDs.
//
// And if you want to help reduce memory thrashing, then the part with the current time should come
// before the part with the randomness.
//
// This will make it so lexical order of the Idempotent IDs is often also a time ordering (for cases
// when the time between 2 IdempotentIDs is not exactly the same).
//
//
// Verb
//
// You can probably think of ‘Verb’ as the name of the instruction.
//
// Some examples Verbs are:
//
// • “empty that shopping cart”,
//
// • “add this book to that shopping cart”,
//
// • “add this item to that TODO list”, and
//
// • “add this e-mail address to my profile”.
//
// Although in the code you may encode these as:
//
// • “EMPTY_SHOPPING_CART”,
//
// • “ADD_TO_SHOPPING_CART”,
//
// • “APPEND”,
//
// • “RECORD_EMAIL”.
//
//
// Verb From HTTP Request
//
// To provide a Verb in an http.Request, have the HTTP client set the
// HTTP request method, or if that is not possible, use X-HTTP-Method-Override
//
// So, consider this example HTTP request:
//
//	POST /v1/user/email HTTP/1.1
//	Host: api.example.com
//	X-Idempotent-ID: z-2015-05-07T10:25:09Z_tleEiguQe67zJFYUa7pngSZT8HX7FMAcHb1Z4yOO2ANtltRPRwF5p9TWwf7m
//	
//	email_address=joeblow@email.com
//
// The Verb here is “POST”.
//
// Now consider this example that includes an “X-HTTP-Method-Override” header:
//
//	POST /v1/user/email HTTP/1.1
//	Host: api.example.com
//	X-Idempotent-ID: z-2015-05-07T10:25:09Z_tleEiguQe67zJFYUa7pngSZT8HX7FMAcHb1Z4yOO2ANtltRPRwF5p9TWwf7m
//	X-HTTP-Method-Override: ADD_EMAIL_ADDRESS
//	
//	email_address=joeblow@email.com
//
// The Verb here is “ADD_EMAIL_ADDRESS”.
//
// Now consider this example that uses a non-regular HTTP method:
//
//	ADD_EMAIL_ADDRESS /v1/user/email HTTP/1.1
//	Host: api.example.com
//	X-Idempotent-ID: z-2015-05-07T10:25:09Z_tleEiguQe67zJFYUa7pngSZT8HX7FMAcHb1Z4yOO2ANtltRPRwF5p9TWwf7m
//	
//	email_address=joeblow@email.com
//
// Again the Verb here is “ADD_EMAIL_ADDRESS”.
//
// If someone is using XMLHttpRequest, to make the HTTP request, then this can be set with code similar to:
//
//	httpRequest = new XMLHttpRequest();
//	
//	// ...
//	
//	httpRequest.open("ADD_EMAIL_ADDRESS", url);
//
// Or:
//
//	httpRequest = new XMLHttpRequest();
//	
//	// ...
//	
//	httpRequest.open("POST", url);
//	httpRequest.setRequestHeader("X-HTTP-Method-Override", "ADD_EMAIL_ADDRESS")
type Instruction struct {

	IdempotentID String
	Verb String
	Data KeyValues
}

// Scan makes ‘mdl.Instruction’ fit the ‘database/sql.Scanner’ interface.
//
// HTTP Request
//
// It also provides an opinionated way of receiving an instruction from an HTTP request.
// I.e., it uses a specific convention for receiving an instruction from an HTTP request.
//
// The instruction ‘Verb’ is given by the HTTP request method. The normal HTTP request method
// can be overridden by the “X-HTTP-Method-Override” header.
//
// The instruction ‘Idempotent ID’ is given by the value of the “X-Idempotent-ID” header.
//
// The instruction ‘data’ is given by the HTTP request body.
//
// Example HTTP Request
//
// Here is an example of .Scan() being used to receive the instruction from the HTTP request.
//
//	func (receiver *MyHandler) ServeHTTP(responseWriter ResponseWriter, request *Request) {
//		
//		// ...
//		
//		var cmd mdl.Instruction
//		
//		// ...
//		
//		err := cmd.Scan(request)
func (receiver *Instruction) Scan(src interface{}) error {
	if nil == receiver {
		return errNilReceiver
	}

	switch casted := src.(type) {
	case *http.Request:

		// verb
		verb, ok := inferVerb(casted)
		if !ok {
			return errBadVerb
		}
		receiver.Verb = SomeString(verb)

		// id
		id, ok := inferID(casted)
		if !ok {
			return errBadID
		}
		receiver.IdempotentID = SomeString(id)

		// data
		if err := inferData(&receiver.Data, casted); nil != err {
			return errBadBody
		}

		return nil
	default:
		return unsupportedSource(src)
	}
}

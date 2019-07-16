package mdl

import (
	"github.com/reiver/go-whitespace"

	"io"
	"unicode/utf8"
)

func keyTokenEncode(writer io.Writer, part string) error {
	if nil == writer {
		return errNilWriter
	}


	var buffer [utf8.UTFMax]byte

	var s string = part
	for 0 < len(s) {
		r, size := utf8.DecodeRuneInString(s)

		switch {
		case '\\' == r:
			buffer[0] = '\\'
			buffer[1] = '\\'

			writer.Write(buffer[:2])
		case '/'  == r:
			buffer[0] = '\\'
			buffer[1] = '/'

			writer.Write(buffer[:2])
		case '{'  == r:
			buffer[0] = '\\'
			buffer[1] = '{'

			writer.Write(buffer[:2])
		case whitespace.IsWhitespace(r):
			buffer[0] = '\\'
			writer.Write(buffer[:1])

			size := utf8.EncodeRune(buffer[:], r)
			writer.Write(buffer[:size])
		default:
			size := utf8.EncodeRune(buffer[:], r)
			writer.Write(buffer[:size])
		}

		before := len(s)
		s = s[size:]
		after := len(s)
		if before <= after {
			return errInternalError
		}
	}

	return nil
}

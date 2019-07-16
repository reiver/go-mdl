package mdl

import (
	"strings"
)

func keyencode(key ...string) string {

	var builder strings.Builder

	for i, token := range key {
		if 0 != i {
			builder.WriteRune('/')
		}

		if err := encodeToken(&builder, token); nil != err {
			return ""
		}
	}

	return builder.String()
}

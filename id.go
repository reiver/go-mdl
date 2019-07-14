package mdl

import (
	"net/http"
)

func inferID(request *http.Request) (string, bool) {
	if nil == request {
		return "", false
	}

	if value := request.Header.Get("X-Idempotent-ID"); "" != value {
		return value, true
	}

	return "", false
}

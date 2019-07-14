package mdl

import (
	"net/http"
)

func inferVerb(request *http.Request) (string, bool) {
	if nil == request {
		return "", false
	}

	switch request.Method {
	case http.MethodPost, http.MethodPatch:
		if value := request.Header.Get("X-HTTP-Method-Override"); "" != value {
			return value, true
		}
	}

	var value string = request.Method
	if "" == value {
		return "", false
	}

	return value, true

}

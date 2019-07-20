package mdl

import (
	"github.com/reiver/go-errhttp"

	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"unsafe"
)

func inferData(keyvalues *KeyValues, request *http.Request) error {
	if nil == keyvalues {
		return errNilKeyValues
	}

	if nil == request {
		return errNilHttpRequest
	}

	contentType := request.Header.Get("Content-Type")

	switch contentType {
	case "application/x-www-form-urlencoded":
		if nil == request.Body {
			return nil
		}

		var limit int64 = 100 << 20 // 100 MB
		var reader io.Reader = io.LimitReader(request.Body, limit)

		var body string
		{
			bytes, err := ioutil.ReadAll(reader)
			if nil != err {
				return errhttp.BadRequestWrap(err)
			}

			{
				var buffer [1]byte
				n, _ := request.Body.Read(buffer[:])
				if n > 0 {
					return errhttp.PayloadTooLargeWrap(fmt.Errorf("HTTP request payload too large; limit = %d MB", limit))
				}
			}

			sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&bytes))
			stringHeader := reflect.StringHeader{Data: sliceHeader.Data, Len: sliceHeader.Len}
			body = *(*string)(unsafe.Pointer(&stringHeader))
		}



		kv, err := url.ParseQuery(body)
		if nil != err {
			return errhttp.BadRequestWrap(err)
		}

		for k, _ := range kv {
			v := kv.Get(k)

			var key Key = SomeKey(k)

			if err := keyvalues.Store(key, v); nil != err {
				return errhttp.BadRequestWrap(err)
			}
		}

		return nil

//	case "application/json":
	default:
		return errhttp.UnsupportedMediaTypeWrap(fmt.Errorf("mdl: %q is an unsupported media type", contentType))
	}
}

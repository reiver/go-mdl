package mdl

import (
	"errors"
)

var (
	errNilReceiver = errors.New("mdl: Nil Receiver")
	errNotLoaded   = errors.New("mdl: Not Loaded")
)

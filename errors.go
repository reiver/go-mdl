package mdl

import (
	"errors"
)

var (
	errEmptyKey       error = internalEmptyKey{}
	errRuneError      error = errors.New("mdl: Rune Error")
)

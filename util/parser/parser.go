package parser

import (
	"strconv"

	"github.com/gustavohmsilva/TechCheck/rendering"
)

// Uint64 will receive a string and generate either a uint64 or a object of
// type rendering.responseError.
func Uint64(s string) (uint64, rendering.ResponseError) {
	if s == "" {
		return 0, rendering.ResponseError{}
	}
	x, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, rendering.ResponseError{Err: err.Error()}

	}
	return x, rendering.ResponseError{}
}

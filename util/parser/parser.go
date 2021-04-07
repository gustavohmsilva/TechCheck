package parser

import (
	"fmt"
	"strconv"
)

// Uint64 will receive a string and generate either a uint64 or a object of
// type rendering.responseError.
func Uint64(s, key string) (uint64, error) {
	if s == "" {
		return 0, nil
	}
	x, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("%v: %v", key, "failed to parse value")

	}
	return x, nil
}

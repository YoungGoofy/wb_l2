package unpack

import (
	"errors"
	"fmt"
	"unicode"
)

// ErrInvalidString custom error
var ErrInvalidString = errors.New("bad string")

// Unpack func unpacking strings
func Unpack(str string) (string, error) {
	rc := []rune(str)
	for index, item := range rc {
		fmt.Println(string(item))
		if unicode.IsDigit(item) && index == 0 {
			return "", ErrInvalidString
		}
	}

	return "", ErrInvalidString
}

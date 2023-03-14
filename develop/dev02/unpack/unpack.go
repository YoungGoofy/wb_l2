package unpack

import (
	"errors"
	"regexp"
	"strings"
	"unicode"
)

// ErrInvalidString custom error
var ErrInvalidString = errors.New("bad string")

// Unpack func unpacking strings
func Unpack(str string) (string, error) {
	var b strings.Builder
	var backslash bool
	var runes rune
	if !CheckLetters(str) {
		return "", nil
	}
	if !CheckFirstItem(str) {
		return "", ErrInvalidString
	}
	for _, char := range str {
		if unicode.IsDigit(char) && !backslash {
			m := int(char - '0')
			r := strings.Repeat(string(runes), m-1)
			b.WriteString(r)
		} else {
			backslash = string(char) == "\\" && string(runes) != "\\"
			if !backslash {
				b.WriteRune(char)
			}
			runes = char
		}
	}
	return b.String(), nil
}

// CheckLetters check if the string is letters
func CheckLetters(str string) bool {
	re := regexp.MustCompile(`[a-zA-Z]+`)
	return re.MatchString(str)
}

// CheckFirstItem check if the string is first item
func CheckFirstItem(str string) bool {
	re := regexp.MustCompile(`^[a-zA-Z]`)
	return re.MatchString(str)
}

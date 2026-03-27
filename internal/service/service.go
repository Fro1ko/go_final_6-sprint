package service

import (
	"errors"
	"strings"
	"unicode"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

var reverseMorse = make(map[string]rune)

func init() {
	for k, v := range morse.DefaultMorse {
		reverseMorse[v] = k
	}
}

func isMorse(s string) bool {
	for _, c := range s {
		if c != '.' && c != '-' && c != ' ' {
			return false
		}
	}
	return true
}
func Convert(s string) (string, error) {
	if s == "" {
		return "", errors.New("input string is empty")
	}

	if isMorse(s) {
		codes := strings.Split(s, " ")
		for _, code := range codes {
			if _, ok := reverseMorse[code]; !ok {
				return "", morse.ErrNoEncoding{Text: code}
			}
		}
		return morse.ToText(s), nil
	} else {
		for _, r := range s {
			r = unicode.ToUpper(r)
			if _, ok := morse.DefaultMorse[r]; !ok {
				return "", morse.ErrNoEncoding{Text: string(r)}
			}
		}
		return morse.ToMorse(s), nil
	}

}

// Package neat implements value sanitisation and conversion functions.
package neat

import (
	"strings"
	"time"
	"unicode"
)

// Body returns a whitespace-trimmed body string with a trailing newline.
func Body(body string) string {
	return strings.TrimSpace(body) + "\n"
}

// Name returns a lowercase alphanumeric-with-dashes name string.
func Name(name string) string {
	var runes []rune
	for _, rune := range strings.ToLower(name) {
		switch {
		case unicode.IsLetter(rune) || unicode.IsNumber(rune):
			runes = append(runes, rune)
		case unicode.IsSpace(rune) || rune == '-':
			runes = append(runes, '-')
		}
	}

	return strings.Trim(string(runes), "-")
}

// Time returns a local Time object from a Unix UTC integer.
func Time(unix int64) time.Time {
	return time.Unix(unix, 0).Local()
}

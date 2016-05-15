package bob // package name must match the package name in bob_test.go
import (
	"strings"
	"unicode"
)

const testVersion = 2 // same as targetTestVersion

const (
	whatever string = "Whatever."
	chillOut        = "Whoa, chill out!"
	sure            = "Sure."
	fine            = "Fine. Be that way!"
)

// Hey returns the correct answer from BOB
func Hey(s string) string {
	s = strings.TrimSpace(s)

	switch {
	case s == "":
		return fine
	case check(s, unicode.IsUpper) && !check(s, unicode.IsLower):
		return chillOut
	case strings.HasSuffix(s, "?"):
		return sure
	default:
		return whatever
	}
}

/*check determines if any items in a string pass some test*/
func check(items string, test func(rune) bool) bool {
	for _, item := range items {
		if test(item) {
			return true
		}
	}
	return false
}

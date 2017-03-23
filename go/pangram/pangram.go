package pangram

import (
	"regexp"
	"strings"
)

const (
	testVersion  = 1
	alphabetSize = 26
)

func IsPangram(text string) bool {
	regex := regexp.MustCompile("[^a-zA-Z]")

	toOnlyValidChars := regex.ReplaceAllString(text, "")
	toLowerCase := strings.ToLower(toOnlyValidChars)
	toArray := strings.Split(toLowerCase, "")
	toSet := createSet(toArray)

	return ifAlphabetSize(toSet)
}

func createSet(letters []string) map[string]bool {
	set := make(map[string]bool)
	for _, r := range letters {
		c := string(r)
		if !set[c] {
			set[c] = true
		}
	}
	return set
}

func ifAlphabetSize(set map[string]bool) bool {
	if len(set) == alphabetSize {
		return true
	}
	return false
}

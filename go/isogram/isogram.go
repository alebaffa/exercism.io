package isogram

import (
	"regexp"
	"strings"
)

const testVersion = 1

func IsIsogram(word string) bool {
	word = removeSpecialChar(word)
	mappa := map[string]bool{}
	for _, letter := range word {
		lowercase := strings.ToLower(string(letter))
		if _, ok := mappa[lowercase]; ok {
			return false
		}
		mappa[lowercase] = true
	}
	return true
}

func removeSpecialChar(word string) string {
	re := regexp.MustCompile("[\\- ]")
	return re.ReplaceAllString(word, "")
}

package isogram

import (
	"unicode"
)

const testVersion = 1

func IsIsogram(word string) bool {
	mappa := make(map[rune]bool)
	for _, letter := range word {
		if !unicode.IsLetter(letter) {
			continue
		}
		letter = unicode.ToLower(letter)
		if mappa[letter] {
			return false
		}
		mappa[letter] = true
	}
	return true
}

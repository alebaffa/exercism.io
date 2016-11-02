package anagram

import (
	"sort"
	"strings"
)

//Detect returns the list of anagram of the input "word" contained in "anagrams"
func Detect(word string, anagrams []string) []string {
	word = strings.ToLower(word)
	result := []string{}

	for _, candidate := range anagrams {
		candidate = strings.ToLower(candidate)

		if haveDifferentSize(word, candidate) {
			continue
		}

		if isAnagram(word, candidate) {
			result = append(result, candidate)
		}

	}
	return result
}

func haveDifferentSize(word, candidate string) bool {
	return len(candidate) != len(word)
}

func isAnagram(word, candidate string) bool {
	return word != candidate && convert(word) == convert(candidate)
}

func convert(s string) string {
	sNoSpace := strings.Split(s, "")
	sort.Strings(sNoSpace)
	return strings.Join(sNoSpace, "")
}

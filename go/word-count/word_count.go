package wordcount

import "regexp"
import "strings"

const testVersion = 2

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	words := getWords(phrase)

	result := make(Frequency)
	for _, word := range words {
		result[strings.ToLower(word)]++
	}
	return result
}

func getWords(phrase string) []string {
	return regexp.MustCompile("[A-Za-z1-9]+").FindAllString(phrase, -1)
}

package wordcount

import "regexp"
import "strings"

const testVersion = 2

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	regex := regexp.MustCompile("[A-Za-z]+|[1-9]+")
	words := regex.FindAllString(phrase, -1)

	result := make(Frequency)
	for _, word := range words {
		result[strings.ToLower(word)]++
	}
	return result
}

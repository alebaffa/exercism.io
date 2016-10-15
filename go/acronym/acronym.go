package acronym

import (
	"fmt"
	"regexp"
	"strings"
)

const testVersion = 1

//abbreviate returns the acronym of a given text
func abbreviate(text string) string {
	regex := regexp.MustCompile("[A-Z]+[a-z]*|[a-z]+")
	words := regex.FindAllString(text, -1)
	acronym := []string{}

	for _, word := range words {
		acronym = append(acronym, string(word[0]))
	}
	return getString(acronym)
}

func getString(acronym []string) string {
	return fmt.Sprintf("%s", strings.ToUpper(strings.Join(acronym, "")))
}

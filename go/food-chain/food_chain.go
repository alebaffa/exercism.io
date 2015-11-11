package foodchain

import (
	"bytes"
	"fmt"
)

const TestVersion = 1

type Phrase struct {
	animal string
	second string
}

var Phrases = []Phrase{
	{"fly", "I don't know why she swallowed the fly. Perhaps she'll die."},
	{"spider", "It wriggled and jiggled and tickled inside her.\n"},
	{"bird", "How absurd to swallow a bird!\n"},
	{"cat", "Imagine that, to swallow a cat!\n"},
	{"dog", "What a hog, to swallow a dog!\n"},
	{"goat", "Just opened her throat and swallowed a goat!\n"},
	{"cow", "I don't know how she swallowed a cow!\n"},
	{"horse", "She's dead, of course!"},
}

func Verse(num int) string {
	num -= 1
	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprint("I know an old lady who swallowed a ", Phrases[num].animal, ".\n"))
	buffer.WriteString(fmt.Sprint(Phrases[num].second))

	if Phrases[num].animal == "fly" || Phrases[num].animal == "horse" {
		return buffer.String()
	}

	for i := num; i >= 1; i-- {
		buffer.WriteString(fmt.Sprint("She swallowed the ", Phrases[i].animal, " to catch the ", Phrases[i-1].animal))
		if Phrases[i-1].animal == "spider" {
			buffer.WriteString(fmt.Sprint(" that wriggled and jiggled and tickled inside her"))
		}
		buffer.WriteString(fmt.Sprint(".\n"))
	}

	buffer.WriteString(fmt.Sprint(Phrases[0].second))
	return buffer.String()
}

func Verses(i, j int) string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprint(Verse(i), "\n", "\n", Verse(j)))
	return buffer.String()

}

func Song() string {
	var buffer bytes.Buffer
	for i := 1; i < 9; i++ {
		if i != 1 {
			buffer.WriteString("\n\n")
		}
		buffer.WriteString(fmt.Sprint(Verse(i)))
	}
	return buffer.String()
}

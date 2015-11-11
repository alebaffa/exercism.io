package foodchain

import (
	"fmt"
	"strings"
)

// TestVersion is the version of this code.
const TestVersion = 1

//Phrase is the structure of the verses composing the poem.
type Phrase struct {
	animal string
	how    string
}

//Phrases is the array of Phrase structures
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

// Verse returns the verse of the poem at the given position.
func Verse(position int) string {
	position--
	poem := fmt.Sprint("I know an old lady who swallowed a ", Phrases[position].animal, ".\n")
	poem += Phrases[position].how

	// if cases of 'fly' and 'horse' just returns now.
	if position == 0 || position == 7 {
		return poem
	}

	for position > 0 {
		poem += fmt.Sprint("She swallowed the ", Phrases[position].animal, " to catch the ", Phrases[position-1].animal)
		if position == 2 {
			poem += " that wriggled and jiggled and tickled inside her.\n"
		} else {
			poem += ".\n"
		}
		position--
	}

	poem += fmt.Sprint(Phrases[0].how)
	return poem
}

//Verses returns the verses of the poem at the given positions.
func Verses(start, end int) string {
	var verses []string
	for i := start; i <= end; i++ {
		verses = append(verses, Verse(i))
	}
	return strings.Join(verses, "\n\n")

}

//Song returns the entire poem.
func Song() string {
	return Verses(1, len(Phrases))
}

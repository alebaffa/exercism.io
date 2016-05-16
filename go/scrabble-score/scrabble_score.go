package scrabble

import "strings"

const testVersion = 4

type Scrabble struct {
	value   int
	letters []string
}

func (s *Scrabble) isContained(input string) bool {
	input = strings.ToLower(input)
	for _, letter := range s.letters {
		if input == string(letter) {
			return true
		}
	}
	return false
}

var scrabbleOne Scrabble
var scrabbleTwo Scrabble
var scrabbleThree Scrabble
var scrabbleFour Scrabble
var scrabbleFive Scrabble
var scrabbleEight Scrabble
var scrabbleTen Scrabble

func init() {
	scrabbleOne = Scrabble{1, []string{"a", "e", "i", "o", "u", "l", "n", "r", "s", "t"}}
	scrabbleTwo = Scrabble{2, []string{"d", "g"}}
	scrabbleThree = Scrabble{3, []string{"b", "c", "m", "p"}}
	scrabbleFour = Scrabble{4, []string{"f", "h", "v", "w", "y"}}
	scrabbleFive = Scrabble{5, []string{"k"}}
	scrabbleEight = Scrabble{8, []string{"j", "x"}}
	scrabbleTen = Scrabble{10, []string{"q", "z"}}

}

func Score(letter string) int {
	score := 0
	if letter == "" {
		return 0
	}

	for _, ch := range letter {
		switch {
		case scrabbleOne.isContained(string(ch)):
			score += scrabbleOne.value
		case scrabbleTwo.isContained(string(ch)):
			score += scrabbleTwo.value
		case scrabbleThree.isContained(string(ch)):
			score += scrabbleThree.value
		case scrabbleFour.isContained(string(ch)):
			score += scrabbleFour.value
		case scrabbleFive.isContained(string(ch)):
			score += scrabbleFive.value
		case scrabbleEight.isContained(string(ch)):
			score += scrabbleEight.value
		case scrabbleTen.isContained(string(ch)):
			score += scrabbleTen.value
		}
	}

	return score
}

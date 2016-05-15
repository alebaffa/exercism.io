package hamming

import (
	"errors"
	"strings"
)

const testVersion = 4

var errSize = errors.New("The two DNA have different size!")

// Distance compares two DNA strands (DNAa, b) and returns
// how many nucleotides are different in the two. */
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errSize
	}

	return countDistance(a, b)
}

func countDistance(a, b string) (int, error) {
	// this check works also with strings that have different
	// lenght of bytes (but still the same num of 'characters' - e.g.:'ä' and 'a')
	if a == b {
		return 0, nil
	}

	distance := 0
	readerA, readerB := getReader(a, b)

	for readerA.Len() != 0 {
		charA, _, _ := readerA.ReadRune()
		charB, _, _ := readerB.ReadRune()
		if charA != charB {
			distance++
		}
	}
	return distance, nil
}

func getReader(a, b string) (readrA, readrB *strings.Reader) {
	return strings.NewReader(a), strings.NewReader(b)
}

package hamming

import "errors"

const testVersion = 4

/*Distance compares two DNA strands (DNAa, b) and returns how many nucleotides are different in the two. */
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("Error")
	}

	return countDistance(a, b)
}

func countDistance(a, b string) (int, error) {
	distance := 0
	for index, value := range a {
		nucleotidA := string(value)
		nucleotidB := string(b[index])

		if nucleotidA != nucleotidB {
			distance++
		}
	}
	return distance, nil
}

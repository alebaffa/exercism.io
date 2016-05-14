package hamming

import "errors"

const testVersion = 4

/*Distance compares two DNA strands (DNAa, b) and returns how many nucleotides are different in the two. */
func Distance(DNAa, DNAb string) (int, error) {
	if len(DNAa) != len(DNAb) {
		return -1, errors.New("Error")
	}
	count := 0
	for indexA, nucleotidA := range DNAa {
		if string(nucleotidA) != string(DNAb[indexA]) {
			count++
		}
	}
	return count, nil
}

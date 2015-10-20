package raindrops

import (
	"strconv"
	"strings"
)

func primeFactors(num int) string {
	//create a slice of string.
	resultString := []string{}
	// append the correct string depending on the prime factor.
	if num%3 == 0 {
		resultString = append(resultString, "Pling")
	}
	if num%5 == 0 {
		resultString = append(resultString, "Plang")
	}
	if num%7 == 0 {
		resultString = append(resultString, "Plong")
	}
	// In case the string so far is empty,
	// it means that prime factors are not 3, 5 nor 7.
	// So, print the original number.
	if len(resultString) == 0 {
		resultString = append(resultString, strconv.Itoa(num))
	}
	return strings.Join(resultString, "")
}

// Convert converts a given int number in to a string
// depending on its prime factors.
func Convert(num int) string {
	return primeFactors(num)
}

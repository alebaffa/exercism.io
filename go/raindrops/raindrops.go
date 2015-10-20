package raindrops

import "math"

//primeFactors returns a string depending on the prime factors of the given input
func primeFactors(num int) string {
	//check if the number is divisible per 2
	for num%2 == 0 {
		num /= 2
	}

	for index := 3; index <= int(math.Sqrt(float64(num))); index += 2 {
		for num%index == 0 {
			n /= index
		}
	}

	if num > 2 {

	}
	return ""
}

//Convert converts a given int number in to a string depending on its prime factors
func Convert(num int) string {
	return primeFactors(num)
}

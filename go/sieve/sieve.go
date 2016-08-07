package sieve

import "math"

//Sieve return list of primes less than N
func Sieve(limit int) (primes []int) {
	b := make([]bool, limit)
	squareLimit := math.Sqrt(float64(limit))
	for i := 2; i <= int(squareLimit); i++ {
		if b[i] == true {
			continue
		}
		primes = append(primes, i)
		for k := i * i; k < limit; k += i {
			b[k] = true
		}
	}
	for i := int(squareLimit) + 1; i < limit; i++ {
		if b[i] == true {
			continue
		}
		primes = append(primes, i)
	}
	return primes
}

package diffsquares

// SquareOfSums returns the square of the sums
func SquareOfSums(n int) int {
	sum := 0
	for num := 1; num <= n; num++ {
		sum += num
	}
	return sum * sum
}

// SumOfSquares returns the sum of the squares
func SumOfSquares(n int) int {
	sum := 0
	for num := 1; num <= n; num++ {
		sum += (num * num)
	}
	return sum
}

// Difference returns the difference between SquareOfSums - SumOfSquares
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}

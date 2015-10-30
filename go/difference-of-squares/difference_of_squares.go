package diffsquares

// SquareOfSums returns the square of the sums
func SquareOfSums(n int) int {
	sum := (n * (n + 1)) / 2
	return sum * sum
}

// SumOfSquares returns the sum of the squares
func SumOfSquares(n int) int {
	sum := n * (2*n + 1) * (n + 1) / 6
	return sum
}

// Difference returns the difference between SquareOfSums - SumOfSquares
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}

package diffsquares

// SquareOfSums returns square of the sum of the first n integers
func SquareOfSums(n int) int {
	var x int
	for i := 1; i <= n; i++ {
		x += i
	}
	return x * x
}

// SumOfSquares returns sum of the squares of the first n integers
func SumOfSquares(n int) int {
	var x int
	for i := 1; i <= n; i++ {
		x += i * i
	}
	return x
}

// Difference returns the difference between the square of the sum of
// the first n integers and the sum of the squares of the first n integers
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}

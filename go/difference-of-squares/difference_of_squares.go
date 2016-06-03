package diffsquares

func sqrt(x int) int {
	return x * x
}

func SquareOfSums(n int) int {
	return sqrt((n * (n + 1)) / 2)
}

func SumOfSquares(n int) int {
	return (n * (n + 1) * (2*n + 1)) / 6
}

func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}

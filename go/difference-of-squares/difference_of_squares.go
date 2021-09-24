package diffsquares

func SquareOfSum(number int) int {
	total := 0
	for i := 1; i <= number; i++ {
		total += i
	}
	return total * total
}

func SumOfSquares(number int) int {
	total := 0
	for i := 1; i <= number; i++ {
		total += (i * i)
	}
	return total
}

func Difference(number int) int {
	squareSum := SquareOfSum(number)
	sumSquares := SumOfSquares(number)
	return squareSum - sumSquares
}

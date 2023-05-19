package Functions

func Factorial[T int | float64](num T) T {
	if int(num) <= 1 {
		return 1
	}
	return num * Factorial(num-1)
}

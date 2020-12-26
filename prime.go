package prime

// FactorsOf returns the prime factors of an integer.
func FactorsOf(n int) []int {
	factors := make([]int, 0)
	for divisor := 2; n > 1; divisor++ {
		for ; n%divisor == 0; n = n / divisor {
			factors = append(factors, divisor)
		}
	}
	return factors
}

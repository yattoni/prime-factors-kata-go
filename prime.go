package prime

import (
	"math"
)

/*
Empty list
*/
func factorsOfOne(n int) []int {
	factors := make([]int, 0)
	return factors
}

/*
Just n == 2 works but n > 1 is more generic

func factorsOfTwo(n int) []int {
	factors := make([]int, 0)
	if n == 2 {
		factors = append(factors, 2)
	}
	return factors
}
*/
func factorsOfTwo(n int) []int {
	factors := make([]int, 0)
	if n > 1 {
		factors = append(factors, 2)
	}
	return factors
}

/*
Change appending 2 to appending n
*/
func factorsOfThree(n int) []int {
	factors := make([]int, 0)
	if n > 1 {
		factors = append(factors, n)
	}
	return factors
}

/*
Second if n > 1 can be inside or outside first one

func factorsOfFour(n int) []int {
	factors := make([]int, 0)
	if n > 1 {
		if n%2 == 0 {
			factors = append(factors, 2)
			n = n / 2
		}
		if n > 1 {
		    factors = append(factors, n)
	    }
	}
	return factors
}

Works for 4,5,6,7
*/
func factorsOfFour(n int) []int {
	factors := make([]int, 0)
	if n > 1 {
		if n%2 == 0 {
			factors = append(factors, 2)
			n = n / 2
		}
	}
	if n > 1 {
		factors = append(factors, n)
	}
	return factors
}

/*
Change if n%2 == 0 to while n%2 == 0
*/
func factorsOfEight(n int) []int {
	factors := make([]int, 0)
	if n > 1 {
		for n%2 == 0 {
			factors = append(factors, 2)
			n = n / 2
		}
	}
	if n > 1 {
		factors = append(factors, n)
	}
	return factors
}

/*
Copy the part that factors out 2s to factor out 3s
*/
func factorsOfNine(n int) []int {
	factors := make([]int, 0)
	if n > 1 {
		for n%2 == 0 {
			factors = append(factors, 2)
			n = n / 2
		}
		for n%3 == 0 {
			factors = append(factors, 3)
			n = n / 3
		}
	}
	if n > 1 {
		factors = append(factors, n)
	}
	return factors
}

/*
Minimal change to get factorsOfN to work.

Deduplicate the two engines that factor out 2,3 to a while loop that factors out a divisor
*/
func factorsOfNMinimal(n int) []int {
	factors := make([]int, 0)
	divisor := 2
	for n > 1 {
		for n%divisor == 0 {
			factors = append(factors, divisor)
			n = n / divisor
		}
		divisor++
	}
	// This if could be removed at this point
	if n > 1 {
		factors = append(factors, n)
	}
	return factors
}

/*
Refactoring

Make inner while loop a for loop

func factorsOfNine(n int) []int {
	factors := make([]int, 0)
	divisor := 2
	for n > 1 {
		for ; n%divisor == 0; n = n / divisor {
			factors = append(factors, divisor)
		}
		divisor++
	}
	if n > 1 {
		factors = append(factors, n)
	}
	return factors
}

Make outer while loop a for loop

func factorsOfNine(n int) []int {
	factors := make([]int, 0)
	for divisor := 2; n > 1; divisor++ {
		for ; n%divisor == 0; n = n / divisor {
			factors = append(factors, divisor)
		}
	}
	if n > 1 {
		factors = append(factors, n)
	}
	return factors
}

Then can delete the trailing if n > 1
*/

func factorsOfN(n int) []int {
	factors := make([]int, 0)
	for divisor := 2; n > 1; divisor++ {
		for ; n%divisor == 0; n = n / divisor {
			factors = append(factors, divisor)
		}
	}
	return factors
}

/*
Termiante outer loop at square root of n.

If we reach the end of the loops and no factors are found yet then the original number must be prime.

Speeds up factorsOfN(2147483647) from 7 seconds to shorter than go test will show me 0.00s.

Speeds up factorsOfN(67280421310721) from over 2 minutes to 0.05 seconds.
*/
func factorsOfNOptimizedForLargePrimes(n int) []int {
	factors := make([]int, 0)
	origN := n
	// Surely there's a better way to do this
	sqrtN := int(math.Ceil(math.Sqrt(float64(n))))
	for divisor := 2; n > 1 && divisor <= sqrtN; divisor++ {
		for ; n%divisor == 0; n = n / divisor {
			factors = append(factors, divisor)
		}
	}
	if len(factors) == 0 && origN != 1 {
		factors = append(factors, origN)
	}
	return factors
}

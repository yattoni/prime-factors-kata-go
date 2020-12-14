package prime

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactorsOfOne(t *testing.T) {
	assert.Equal(t, []int{}, factorsOfOne(1), "Factors of 1")
}

func TestFactorsOfTwo(t *testing.T) {
	assert.Equal(t, []int{}, factorsOfTwo(1), "Factors of 1")
	assert.Equal(t, []int{2}, factorsOfTwo(2), "Factors of 2")
}

func TestFactorsOfThree(t *testing.T) {
	assert.Equal(t, []int{}, factorsOfThree(1), "Factors of 1")
	assert.Equal(t, []int{2}, factorsOfThree(2), "Factors of 2")
	assert.Equal(t, []int{3}, factorsOfThree(3), "Factors of 3")
}

func TestFactorsOfFour(t *testing.T) {
	assert.Equal(t, []int{}, factorsOfFour(1), "Factors of 1")
	assert.Equal(t, []int{2}, factorsOfFour(2), "Factors of 2")
	assert.Equal(t, []int{3}, factorsOfFour(3), "Factors of 3")
	assert.Equal(t, []int{2, 2}, factorsOfFour(4), "Factors of 4")
}

func TestFactorsOfFiveSixSeven(t *testing.T) {
	// works with factorsOfFour
	assert.Equal(t, []int{}, factorsOfFour(1), "Factors of 1")
	assert.Equal(t, []int{2}, factorsOfFour(2), "Factors of 2")
	assert.Equal(t, []int{3}, factorsOfFour(3), "Factors of 3")
	assert.Equal(t, []int{2, 2}, factorsOfFour(4), "Factors of 4")
	assert.Equal(t, []int{5}, factorsOfFour(5), "Factors of 5")
	assert.Equal(t, []int{2, 3}, factorsOfFour(6), "Factors of 6")
	assert.Equal(t, []int{7}, factorsOfFour(7), "Factors of 7")
}

func TestFactorsOfEight(t *testing.T) {
	assert.Equal(t, []int{}, factorsOfEight(1), "Factors of 1")
	assert.Equal(t, []int{2}, factorsOfEight(2), "Factors of 2")
	assert.Equal(t, []int{3}, factorsOfEight(3), "Factors of 3")
	assert.Equal(t, []int{2, 2}, factorsOfEight(4), "Factors of 4")
	assert.Equal(t, []int{5}, factorsOfEight(5), "Factors of 5")
	assert.Equal(t, []int{2, 3}, factorsOfEight(6), "Factors of 6")
	assert.Equal(t, []int{7}, factorsOfEight(7), "Factors of 7")
	assert.Equal(t, []int{2, 2, 2}, factorsOfEight(8), "Factors of 8")
}

func TestFactorsOfNine(t *testing.T) {
	assert.Equal(t, []int{}, factorsOfNine(1), "Factors of 1")
	assert.Equal(t, []int{2}, factorsOfNine(2), "Factors of 2")
	assert.Equal(t, []int{3}, factorsOfNine(3), "Factors of 3")
	assert.Equal(t, []int{2, 2}, factorsOfNine(4), "Factors of 4")
	assert.Equal(t, []int{5}, factorsOfNine(5), "Factors of 5")
	assert.Equal(t, []int{2, 3}, factorsOfNine(6), "Factors of 6")
	assert.Equal(t, []int{7}, factorsOfNine(7), "Factors of 7")
	assert.Equal(t, []int{2, 2, 2}, factorsOfNine(8), "Factors of 8")
	assert.Equal(t, []int{3, 3}, factorsOfNine(9), "Factors of 9")
}

func TestFactorsOfNMinimal(t *testing.T) {
	assert.Equal(t, []int{}, factorsOfNMinimal(1), "Factors of 1")
	assert.Equal(t, []int{2}, factorsOfNMinimal(2), "Factors of 2")
	assert.Equal(t, []int{3}, factorsOfNMinimal(3), "Factors of 3")
	assert.Equal(t, []int{2, 2}, factorsOfNMinimal(4), "Factors of 4")
	assert.Equal(t, []int{5}, factorsOfNMinimal(5), "Factors of 5")
	assert.Equal(t, []int{2, 3}, factorsOfNMinimal(6), "Factors of 6")
	assert.Equal(t, []int{7}, factorsOfNMinimal(7), "Factors of 7")
	assert.Equal(t, []int{2, 2, 2}, factorsOfNMinimal(8), "Factors of 8")
	assert.Equal(t, []int{3, 3}, factorsOfNMinimal(9), "Factors of 9")
	assert.Equal(t, []int{2, 2, 3, 3, 5, 7, 11, 11, 13}, factorsOfNMinimal(2*2*3*3*5*7*11*11*13), "Factors of 2 * 2 * 3 * 3 * 5 * 7 * 11 * 11 * 13")
}

func TestFactorsOfN(t *testing.T) {
	assert.Equal(t, []int{}, factorsOfN(1), "Factors of 1")
	assert.Equal(t, []int{2}, factorsOfN(2), "Factors of 2")
	assert.Equal(t, []int{3}, factorsOfN(3), "Factors of 3")
	assert.Equal(t, []int{2, 2}, factorsOfN(4), "Factors of 4")
	assert.Equal(t, []int{5}, factorsOfN(5), "Factors of 5")
	assert.Equal(t, []int{2, 3}, factorsOfN(6), "Factors of 6")
	assert.Equal(t, []int{7}, factorsOfN(7), "Factors of 7")
	assert.Equal(t, []int{2, 2, 2}, factorsOfN(8), "Factors of 8")
	assert.Equal(t, []int{3, 3}, factorsOfN(9), "Factors of 9")
	assert.Equal(t, []int{2, 2, 3, 3, 5, 7, 11, 11, 13}, factorsOfN(2*2*3*3*5*7*11*11*13), "Factors of 2 * 2 * 3 * 3 * 5 * 7 * 11 * 11 * 13")
}

/*
https://en.wikipedia.org/wiki/2,147,483,647
*/
func TestFactorsOfN_EighthMersennePrime(t *testing.T) {
	t.Skip("Takes ~7 seconds")
	assert.Equal(t, []int{2, 3, 3, 7, 11, 31, 151, 331}, factorsOfN(2147483647), "Factors of 2,147,483,647")
}

func TestFactorsOfN_EighthMersennePrimeMinusOne(t *testing.T) {
	assert.Equal(t, []int{2, 3, 3, 7, 11, 31, 151, 331}, factorsOfN(2147483647-1), "Factors of 2,147,483,647")
}

/*
https://en.wikipedia.org/wiki/Largest_known_prime_number
*/
func TestFactorsOfN_LargePrimeNumbers(t *testing.T) {
	t.Skip("Takes at least 2 minutes")
	assert.Equal(t, []int{67280421310721}, factorsOfN(67280421310721), "Factors of 67,280,421,310,721")
}

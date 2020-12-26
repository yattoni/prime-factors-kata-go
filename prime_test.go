package prime

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactorsOf(t *testing.T) {
	assert.Equal(t, []int{}, FactorsOf(1))
	assert.Equal(t, []int{2}, FactorsOf(2))
	assert.Equal(t, []int{3}, FactorsOf(3))
	assert.Equal(t, []int{2, 2}, FactorsOf(4))
	assert.Equal(t, []int{5}, FactorsOf(5))
	assert.Equal(t, []int{2, 3}, FactorsOf(6))
	assert.Equal(t, []int{7}, FactorsOf(7))
	assert.Equal(t, []int{2, 2, 2}, FactorsOf(8))
	assert.Equal(t, []int{3, 3}, FactorsOf(9))
	assert.Equal(t, []int{2, 2, 3, 3, 5, 7, 11, 11, 13}, FactorsOf(2*2*3*3*5*7*11*11*13))
}

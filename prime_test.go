package prime

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactorsOf(t *testing.T) {
	assert.Equal(t, []int{}, FactorsOf(1))
}

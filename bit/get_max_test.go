package bit

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMax(t *testing.T) {
	a := -16
	b := 1

	assert.Equal(t, b, GetMax(a, b))

	a = 2147483647
	b = -2147480000

	assert.Equal(t, a, GetMax(a, b))
}

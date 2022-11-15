package dynamic_programming

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLivePossibility(t *testing.T) {
	ans1 := LivePossibility1(6, 6, 10, 50, 50)
	ans2 := LivePossibility2(6, 6, 10, 50, 50)

	assert.Equal(t, ans1, ans2)
}

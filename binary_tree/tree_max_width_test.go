package binary_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxWidth(t *testing.T) {
	testTimes := 1000000
	maxLevel := 10
	maxValue := 100

	for i := 0; i < testTimes; i++ {
		head := GenerateRandomBST(maxLevel, maxValue)

		w1 := MaxWidthUseMap(head)
		w2 := MaxWidthNoMap(head)

		if !assert.Equal(t, w1, w2) {
			break
		}
	}
}

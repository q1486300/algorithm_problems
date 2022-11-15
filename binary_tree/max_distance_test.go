package binary_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxDistance(t *testing.T) {
	testTimes := 1000000
	maxLevel := 5
	maxValue := 100

	for i := 0; i < testTimes; i++ {
		head := GenerateRandomBST(maxLevel, maxValue)

		d1 := MaxDistance1(head)
		d2 := MaxDistance2(head)

		if !assert.Equal(t, d2, d1) {
			break
		}
	}
}

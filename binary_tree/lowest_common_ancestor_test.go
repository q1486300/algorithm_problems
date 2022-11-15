package binary_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLowestCommonAncestor(t *testing.T) {
	testTimes := 1000000
	maxLevel := 5
	maxValue := 100

	for i := 0; i < testTimes; i++ {
		head := GenerateRandomBST(maxLevel, maxValue)

		o1 := PickRandomOne(head)
		o2 := PickRandomOne(head)

		n1 := LowestCommonAncestor1(head, o1, o2)
		n2 := LowestCommonAncestor2(head, o1, o2)

		if !assert.Equal(t, n1, n2) {
			break
		}
	}
}

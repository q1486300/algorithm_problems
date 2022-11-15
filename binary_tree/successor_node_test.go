package binary_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetSuccessorNode(t *testing.T) {
	testTimes := 1000000
	maxLevel := 5
	maxValue := 100

	for i := 0; i < testTimes; i++ {
		head := GenerateRandomBSTWithParent(maxLevel, maxValue)

		cur, post1 := PickRandomOneInListAndTheNext(head)
		post2 := GetSuccessorNode(cur)

		if !assert.Equal(t, post1, post2) {
			break
		}
	}
}

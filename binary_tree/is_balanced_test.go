package binary_tree

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func TestIsBalanced(t *testing.T) {
	testTimes := 1000000
	maxLevel := 5
	maxValue := 100

	for i := 0; i < testTimes; i++ {
		head := GenerateRandomBST(maxLevel, maxValue)

		b1 := IsBalanced1(head)
		b2 := IsBalanced2(head)

		if !assert.Equal(t, b1, b2) {
			break
		}
	}
}
func init() {
	rand.Seed(time.Now().UnixNano())
}

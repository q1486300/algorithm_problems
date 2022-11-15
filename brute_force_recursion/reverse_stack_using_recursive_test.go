package brute_force_recursion

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverse(t *testing.T) {
	test := [][]int{
		{1, 2, 3, 4, 5},
		{5, 4, 3, 2, 1},
		{1, 3, 5, 7, 9},
		{2, 4, 6, 8, 10},
		{10, 40, 20, 50, 30},
	}

	for _, ints := range test {
		copyInts := make([]int, len(ints))
		copy(copyInts, ints)

		Reverse(&ints)

		for i, copyInt := range copyInts {
			if !assert.Equal(t, copyInt, ints[len(ints)-1-i]) {
				break
			}
		}
	}
}

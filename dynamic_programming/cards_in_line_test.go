package dynamic_programming

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestWin(t *testing.T) {
	testTimes := 1000000
	maxLen := 15
	maxValue := 20

	for i := 0; i < testTimes; i++ {
		arr := winRandomArray(maxLen, maxValue)

		ans1 := Win1(arr)
		ans2 := Win2(arr)

		if !assert.Equal(t, ans1, ans2) {
			break
		}
	}
}

func winRandomArray(maxLen, maxValue int) []int {
	N := rand.Intn(maxLen)
	arr := make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = rand.Intn(maxValue) + 1
	}
	return arr
}

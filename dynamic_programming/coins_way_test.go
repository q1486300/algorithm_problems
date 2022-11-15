package dynamic_programming

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestCoinsWay(t *testing.T) {
	testTimes := 1000000
	maxLen := 10
	maxValue := 30

	for i := 0; i < testTimes; i++ {
		arr := coinsWayRandomArray(maxLen, maxValue)
		aim := rand.Intn(maxValue)

		ans1 := CoinsWay(arr, aim)
		ans2 := CoinsWayDP1(arr, aim)
		ans3 := CoinsWayDP2(arr, aim)

		if !assert.Equal(t, ans1, ans2) || !assert.Equal(t, ans2, ans3) {
			break
		}
	}
}

func coinsWayRandomArray(maxLen, maxValue int) []int {
	N := rand.Intn(maxLen)
	arr := make([]int, N)
	has := make([]bool, maxValue+1)
	for i := 0; i < N; i++ {
		for {
			arr[i] = rand.Intn(maxValue) + 1
			if !has[arr[i]] {
				break
			}
		}
		has[arr[i]] = true
	}
	return arr
}

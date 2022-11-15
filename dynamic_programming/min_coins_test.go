package dynamic_programming

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestMinCoins(t *testing.T) {
	testTimes := 300000
	maxLen := 20
	maxValue := 30

	for i := 0; i < testTimes; i++ {
		arr := randomArray(maxLen, maxValue)
		aim := rand.Intn(maxValue)

		ans1 := MinCoins(arr, aim)
		ans2 := MinCoinsDP1(arr, aim)
		ans3 := MinCoinsDP2(arr, aim)

		if !assert.Equal(t, ans1, ans2) {
			break
		}
		if !assert.Equal(t, ans2, ans3) {
			break
		}
	}
}

func randomArray(maxLen, maxValue int) []int {
	N := rand.Intn(maxLen)
	arr := make([]int, N)
	has := make([]bool, maxValue+1)
	for i := 0; i < N; i++ {
		for {
			arr[i] = rand.Intn(maxValue) + 1
			if has[arr[i]] != true {
				break
			}
		}
		has[arr[i]] = true
	}
	return arr
}

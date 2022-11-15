package sliding_window

import (
	"github.com/stretchr/testify/assert"
	"math"
	"math/rand"
	"testing"
)

func TestGetMaxWindow(t *testing.T) {
	testTimes := 100000
	maxSize := 100
	maxValue := 100

	for i := 0; i < testTimes; i++ {
		arr := generateRandomArray(maxSize, maxValue)
		w := rand.Intn(len(arr) + 1)

		ans1 := GetMaxWindow(arr, w)
		ans2 := right(arr, w)

		if !assert.Equal(t, ans2, ans1) {
			break
		}
	}
}

func right(arr []int, w int) []int {
	if arr == nil || w < 1 || len(arr) < w {
		return nil
	}
	N := len(arr)
	res := make([]int, N-w+1)
	index := 0
	L := 0
	R := w - 1
	for R < N {
		max := arr[L]
		for i := L + 1; i <= R; i++ {
			max = int(math.Max(float64(max), float64(arr[i])))
		}
		res[index] = max
		index++
		L++
		R++
	}
	return res
}

func generateRandomArray(maxSize, maxValue int) []int {
	arr := make([]int, rand.Intn(maxSize+1))
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(maxValue + 1)
	}
	return arr
}

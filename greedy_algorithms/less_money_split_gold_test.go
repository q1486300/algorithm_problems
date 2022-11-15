package greedy_algorithms

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestLessMoney(t *testing.T) {
	testTimes := 100000
	maxSize := 6
	maxValue := 1000

	for i := 0; i < testTimes; i++ {
		arr := generateRandomArray(maxSize, maxValue)

		m1 := LessMoney1(arr)
		m2 := LessMoney2(arr)

		if !assert.Equal(t, m1, m2) {
			break
		}
	}
}

func generateRandomArray(maxSize, maxValue int) []int {
	arr := make([]int, rand.Intn(maxSize+1))
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(maxValue + 1)
	}
	return arr
}

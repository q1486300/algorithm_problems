package monotonous_stack

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func TestGetNearLess(t *testing.T) {
	testTimes := 2000000
	size := 10
	max := 20

	for i := 0; i < testTimes; i++ {
		arr1 := getRandomArrayNoRepeat(size)
		arr2 := getRandomArray(size, max)

		ans1 := GetNearLessNoRepeat(arr1)
		ans2 := GetNearLess(arr2)

		ans3 := right(arr1)
		ans4 := right(arr2)

		if !assert.Equal(t, ans3, ans1) {
			break
		}

		if !assert.Equal(t, ans4, ans2) {
			break
		}
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func right(arr []int) [][]int {
	res := make([][]int, len(arr))
	for i := range res {
		res[i] = make([]int, 2)
	}

	for i := 0; i < len(arr); i++ {
		leftLessIndex := -1
		rightLessIndex := -1
		cur := i - 1
		for cur >= 0 {
			if arr[cur] < arr[i] {
				leftLessIndex = cur
				break
			}
			cur--
		}
		cur = i + 1
		for cur < len(arr) {
			if arr[cur] < arr[i] {
				rightLessIndex = cur
				break
			}
			cur++
		}
		res[i][0] = leftLessIndex
		res[i][1] = rightLessIndex
	}
	return res
}

func getRandomArrayNoRepeat(size int) []int {
	arr := make([]int, rand.Intn(size)+1)
	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}
	for i := 0; i < len(arr); i++ {
		swapIndex := rand.Intn(len(arr))
		tmp := arr[swapIndex]
		arr[swapIndex] = arr[i]
		arr[i] = tmp
	}
	return arr
}

func getRandomArray(size, max int) []int {
	arr := make([]int, rand.Intn(size)+1)
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(max) - rand.Intn(max)
	}
	return arr
}

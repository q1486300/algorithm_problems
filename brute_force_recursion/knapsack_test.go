package brute_force_recursion

import (
	"fmt"
	"testing"
)

func TestMaxValue(t *testing.T) {
	weights := []int{3, 2, 4, 7, 3, 1, 7}
	values := []int{5, 6, 3, 19, 12, 4, 2}
	bag := 15

	fmt.Println(MaxValue(weights, values, bag))
}

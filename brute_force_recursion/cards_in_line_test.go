package brute_force_recursion

import (
	"fmt"
	"testing"
)

func TestWin(t *testing.T) {
	arr := []int{5, 7, 4, 5, 8, 1, 6, 0, 3, 4, 6, 1, 7}
	fmt.Println(Win(arr))
}

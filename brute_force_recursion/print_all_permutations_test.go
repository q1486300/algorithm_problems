package brute_force_recursion

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrintAllPermutation(t *testing.T) {
	test := "acc"
	ans1 := PrintAllPermutation1(test)
	ans2 := PrintAllPermutation2(test)
	ans3 := PrintAllPermutation3(test)

	assert.ElementsMatch(t, ans1, ans2)

	for _, cur := range ans3 {
		fmt.Println(cur)
	}
}

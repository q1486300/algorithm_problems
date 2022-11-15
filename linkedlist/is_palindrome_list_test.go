package linkedlist

import (
	"fmt"
	"testing"
)

var heads []*Node
var nodeValues = [][]int{
	nil,
	{1},
	{1, 2},
	{1, 1},
	{1, 2, 3},
	{1, 2, 1},
	{1, 2, 3, 1},
	{1, 2, 2, 1},
	{1, 2, 3, 2, 1},
}

func TestIsPalindrome(t *testing.T) {
	for _, head := range heads {
		printLinkedList(head)
		t.Logf("%v | %v | %v\n", IsPalindrome1(head), IsPalindrome2(head), IsPalindrome3(head))
		printLinkedList(head)
		fmt.Println("=============================")
	}
}

func init() {
	for _, value := range nodeValues {
		heads = append(heads, generateLinkedListWithSlice(value))
	}
}

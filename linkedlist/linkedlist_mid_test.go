package linkedlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedListMid(t *testing.T) {
	var heads []*Node
	var nodeValues = [][]int{
		{0, 1},
		{0, 1, 2},
		{0, 1, 2, 3},
		{0, 1, 2, 3, 4},
		{0, 1, 2, 3, 4, 5},
		{0, 1, 2, 3, 4, 5, 6},
		{0, 1, 2, 3, 4, 5, 6, 7},
		{0, 1, 2, 3, 4, 5, 6, 7, 8},
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	}
	generateHeads(&heads, nodeValues)

	for _, head := range heads {
		printLinkedList(head)

		ans1 := MidOrUpMidNode(head)
		ans2 := midOrUpMidNodeComparator(head)
		t.Log(ans1, ans2)
		if !assert.Equal(t, ans2, ans1) {
			break
		}

		ans1 = MidOrDownMidNode(head)
		ans2 = midOrDownMidNodeComparator(head)
		t.Log(ans1, ans2)
		if !assert.Equal(t, ans2, ans1) {
			break
		}

		ans1 = MidOrUpMidPreNode(head)
		ans2 = midOrUpMidPreNodeComparator(head)
		t.Log(ans1, ans2)
		if !assert.Equal(t, ans2, ans1) {
			break
		}

		ans1 = MidOrDownMidPreNode(head)
		ans2 = midOrDownMidPreNodeComparator(head)
		t.Log(ans1, ans2)
		if !assert.Equal(t, ans2, ans1) {
			break
		}
	}
}

func generateHeads(heads *[]*Node, values [][]int) {
	for _, value := range values {
		*heads = append(*heads, generateLinkedListWithSlice(value))
	}
}

func midOrUpMidNodeComparator(head *Node) *Node {
	if head == nil {
		return nil
	}
	cur := head
	var arr []*Node
	for cur != nil {
		arr = append(arr, cur)
		cur = cur.next
	}
	return arr[(len(arr)-1)/2]
}

func midOrDownMidNodeComparator(head *Node) *Node {
	if head == nil {
		return nil
	}
	cur := head
	var arr []*Node
	for cur != nil {
		arr = append(arr, cur)
		cur = cur.next
	}
	return arr[len(arr)/2]
}

func midOrUpMidPreNodeComparator(head *Node) *Node {
	if head == nil || head.next == nil || head.next.next == nil {
		return nil
	}
	cur := head
	var arr []*Node
	for cur != nil {
		arr = append(arr, cur)
		cur = cur.next
	}
	return arr[(len(arr)-3)/2]
}

func midOrDownMidPreNodeComparator(head *Node) *Node {
	if head == nil || head.next == nil {
		return nil
	}
	cur := head
	var arr []*Node
	for cur != nil {
		arr = append(arr, cur)
		cur = cur.next
	}
	return arr[(len(arr)-2)/2]
}

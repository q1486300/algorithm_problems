package linkedlist

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestReverseLinkedList(t *testing.T) {
	length := 50
	value := 100
	testTime := 100000

	for i := 0; i < testTime; i++ {
		// LinkedList
		node1 := generateRandomLinkedList(length, value)
		list1 := getLinkedListOriginOrder(node1)
		node1 = ReverseLinkedList(node1)
		if !assert.True(t, checkLinkedListReverse(list1, node1)) {
			break
		}

		node2 := generateRandomLinkedList(length, value)
		list2 := getLinkedListOriginOrder(node2)
		node2 = reverseLinkedListComparator(node2)
		if !assert.True(t, checkLinkedListReverse(list2, node2)) {
			break
		}

		// Double LinkedList
		node3 := generateRandomDoubleList(length, value)
		list3 := getDoubleListOriginOrder(node3)
		node3 = ReverseDoubleList(node3)
		if !assert.True(t, checkDoubleListReverse(list3, node3)) {
			break
		}

		node4 := generateRandomDoubleList(length, value)
		list4 := getDoubleListOriginOrder(node4)
		node4 = reverseDoubleListComparator(node4)
		if !assert.True(t, checkDoubleListReverse(list4, node4)) {
			break
		}
	}
}

func reverseLinkedListComparator(head *Node) *Node {
	if head == nil {
		return nil
	}
	var list []*Node
	for head != nil {
		list = append(list, head)
		head = head.next
	}
	list[0].next = nil
	for i := 1; i < len(list); i++ {
		list[i].next = list[i-1]
	}
	return list[len(list)-1]
}

func reverseDoubleListComparator(head *DoubleNode) *DoubleNode {
	if head == nil {
		return nil
	}
	var list []*DoubleNode
	for head != nil {
		list = append(list, head)
		head = head.next
	}
	list[0].next = nil
	pre := list[0]
	for i := 1; i < len(list); i++ {
		cur := list[i]
		cur.last = nil
		cur.next = pre
		pre.last = cur
		pre = cur
	}
	return list[len(list)-1]
}

func generateRandomLinkedList(len, value int) *Node {
	size := rand.Intn(len + 1)
	if size == 0 {
		return nil
	}
	size--
	head := NewNode(rand.Intn(value + 1))
	pre := head
	for size != 0 {
		cur := NewNode(rand.Intn(value + 1))
		pre.next = cur
		pre = cur
		size--
	}
	return head
}

func generateRandomDoubleList(len, value int) *DoubleNode {
	size := rand.Intn(len + 1)
	if size == 0 {
		return nil
	}
	size--
	head := NewDoubleNode(rand.Intn(value + 1))
	pre := head
	for size != 0 {
		cur := NewDoubleNode(rand.Intn(value + 1))
		pre.next = cur
		cur.last = pre
		pre = cur
		size--
	}
	return head
}

func getLinkedListOriginOrder(head *Node) []int {
	var ans []int
	for head != nil {
		ans = append(ans, head.value)
		head = head.next
	}
	return ans
}

func checkLinkedListReverse(origin []int, head *Node) bool {
	for i := len(origin) - 1; i >= 0; i-- {
		if origin[i] != head.value {
			return false
		}
		head = head.next
	}
	return true
}

func getDoubleListOriginOrder(head *DoubleNode) []int {
	var ans []int
	for head != nil {
		ans = append(ans, head.value)
		head = head.next
	}
	return ans
}

func checkDoubleListReverse(origin []int, head *DoubleNode) bool {
	var end *DoubleNode
	for i := len(origin) - 1; i >= 0; i-- {
		if origin[i] != head.value {
			return false
		}
		end = head
		head = head.next
	}
	for i := 0; i < len(origin); i++ {
		if origin[i] != end.value {
			return false
		}
		end = end.last
	}
	return true
}

package linkedlist

import (
	"testing"
)

func TestPrintCommonPart(t *testing.T) {
	length := 50
	value := 10
	testTime := 10

	for i := 0; i < testTime; i++ {
		node1 := generateRandomLinkedList(length, value)
		node2 := generateRandomLinkedList(length, value)

		printLinkedList(node1)
		printLinkedList(node2)
		PrintCommonPart(node1, node2)
	}
}

package linkedlist

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestListPartition(t *testing.T) {
	testTime := 500000
	length := 50
	value := 50

	for i := 0; i < testTime; i++ {
		// value:  0 ~ 50
		head := generateRandomLinkedList(length, value)
		// partitionValue: 0 ~ 50
		partitionValue := rand.Intn(value + 1)
		head = ListPartition1(head, partitionValue)
		if !assert.True(t, checkPartition(head, partitionValue)) {
			break
		}

		head = generateRandomLinkedList(length, value)
		partitionValue = rand.Intn(value + 1)
		head = ListPartition2(head, partitionValue)
		if !assert.True(t, checkPartition(head, partitionValue)) {
			break
		}
	}
}

func checkPartition(head *Node, pivot int) bool {
	if head == nil || head.next == nil {
		return true
	}
	var equalFlag, biggerFlag bool
	cur := head
	for cur != nil {
		if cur.value < pivot {
			if equalFlag || biggerFlag {
				return false
			}
		} else if cur.value == pivot {
			equalFlag = true
			if biggerFlag {
				return false
			}
		} else {
			biggerFlag = true
		}
		cur = cur.next
	}
	return true
}

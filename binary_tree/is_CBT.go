package binary_tree

import (
	"container/list"
	"math"
)

func IsCBT1(head *Node) bool {
	if head == nil {
		return true
	}
	queue := list.New()
	leaf := false // 是否遇到過左右孩子不雙全的節點
	var l, r *Node

	queue.PushBack(head)
	for queue.Len() != 0 {
		head = queue.Front().Value.(*Node)
		queue.Remove(queue.Front())
		l = head.left
		r = head.right
		// 如果遇到了不雙全的節點之後，又發現目前節點不是葉節點
		if (leaf && (l != nil || r != nil)) ||
			// 目前節點只有右孩子沒有左孩子，不是完全二元樹
			(l == nil && r != nil) {
			return false
		}
		if l != nil {
			queue.PushBack(l)
		}
		if r != nil {
			queue.PushBack(r)
		}
		// 遇到了不雙全的節點
		if l == nil || r == nil {
			leaf = true
		}
	}
	return true
}

func IsCBT2(head *Node) bool {
	if head == nil {
		return true
	}
	return processIsCBT(head).isCBT
}

func processIsCBT(head *Node) IsCBTInfo {
	if head == nil {
		return NewIsCBTInfo(true, true, 0)
	}
	leftInfo := processIsCBT(head.left)
	rightInfo := processIsCBT(head.right)

	height := int(math.Max(float64(leftInfo.height), float64(rightInfo.height))) + 1

	isFull := leftInfo.isFull && rightInfo.isFull && leftInfo.height == rightInfo.height

	isCBT := false
	if isFull {
		isCBT = true
	} else {
		if leftInfo.isCBT && rightInfo.isCBT {
			if leftInfo.isCBT && rightInfo.isFull && leftInfo.height == rightInfo.height+1 {
				isCBT = true
			}
			if leftInfo.isFull && rightInfo.isFull && leftInfo.height == rightInfo.height+1 {
				isCBT = true
			}
			if leftInfo.isFull && rightInfo.isCBT && leftInfo.height == rightInfo.height {
				isCBT = true
			}
		}
	}
	return NewIsCBTInfo(isFull, isCBT, height)
}

type IsCBTInfo struct {
	isFull, isCBT bool
	height        int
}

func NewIsCBTInfo(isFull, isCBT bool, height int) IsCBTInfo {
	return IsCBTInfo{
		isFull: isFull,
		isCBT:  isCBT,
		height: height,
	}
}

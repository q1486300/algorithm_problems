package binary_tree

import (
	"fmt"
	"math"
)

func Morris(head *Node) {
	if head == nil {
		return
	}
	cur := head
	var mostRight *Node
	for cur != nil {
		mostRight = cur.left
		if mostRight != nil {
			for mostRight.right != nil && mostRight.right != cur {
				mostRight = mostRight.right
			}
			if mostRight.right == nil {
				mostRight.right = cur
				cur = cur.left
				continue
			} else {
				mostRight.right = nil
			}
		}
		cur = cur.right
	}
}

func MorrisPreOrder(head *Node) {
	if head == nil {
		return
	}
	cur := head
	var mostRight *Node
	for cur != nil {
		mostRight = cur.left
		if mostRight != nil {
			for mostRight.right != nil && mostRight.right != cur {
				mostRight = mostRight.right
			}
			if mostRight.right == nil {
				fmt.Printf("%d ", cur.value) // 第一次來到能到達自己兩次的節點
				mostRight.right = cur
				cur = cur.left
				continue
			} else {
				mostRight.right = nil
			}
		} else {
			fmt.Printf("%d ", cur.value) // 第一次來到只能到達自己一次的節點(沒有左子樹)
		}
		cur = cur.right
	}
	fmt.Println()
}

func MorrisInOrder(head *Node) {
	if head == nil {
		return
	}
	cur := head
	var mostRight *Node
	for cur != nil {
		mostRight = cur.left
		if mostRight != nil {
			for mostRight.right != nil && mostRight.right != cur {
				mostRight = mostRight.right
			}
			if mostRight.right == nil {
				mostRight.right = cur
				cur = cur.left
				continue
			} else {
				mostRight.right = nil
			}
		}
		// 第一次來到只能到達自己一次的節點(沒有左子樹)
		// 第二次來到能到達自己兩次的節點
		fmt.Printf("%d ", cur.value)
		cur = cur.right
	}
	fmt.Println()
}

func MorrisPostOrder(head *Node) {
	if head == nil {
		return
	}
	cur := head
	var mostRight *Node
	for cur != nil {
		mostRight = cur.left
		if mostRight != nil {
			for mostRight.right != nil && mostRight.right != cur {
				mostRight = mostRight.right
			}
			if mostRight.right == nil {
				mostRight.right = cur
				cur = cur.left
				continue
			} else {
				mostRight.right = nil
				printEdge(cur.left) // 第二次來到能到達自己兩次的節點時，印出左子樹右邊界的逆序
			}
		}
		cur = cur.right
	}
	printEdge(head) // 遍歷完所有節點時，印出根節點右邊界的逆序
	fmt.Println()
}

func printEdge(head *Node) {
	tail := reverseEdge(head)
	cur := tail
	for cur != nil {
		fmt.Printf("%d ", cur.value)
		cur = cur.right
	}
	reverseEdge(tail)
}

func reverseEdge(from *Node) *Node {
	var pre, next *Node
	for from != nil {
		next = from.right
		from.right = pre
		pre = from
		from = next
	}
	return pre
}

func MorrisIsBST(head *Node) bool {
	if head == nil {
		return true
	}
	cur := head
	var mostRight *Node
	atLeastOneNode := false
	preValue := math.MinInt // 中序前一個節點的值
	for cur != nil {
		mostRight = cur.left
		if mostRight != nil {
			for mostRight.right != nil && mostRight.right != cur {
				mostRight = mostRight.right
			}
			if mostRight.right == nil {
				mostRight.right = cur
				cur = cur.left
				continue
			} else {
				mostRight.right = nil
			}
		}
		// atLeastOneNode 避免中序遍歷的第一個值是 math.MinInt
		if atLeastOneNode && cur.value <= preValue { // 目前節點 <= 前一個節點的值，返回 false
			return false
		}
		atLeastOneNode = true
		preValue = cur.value // 更新前一個節點的值為目前節點
		cur = cur.right
	}
	return true // 所有節點成功遍歷完，返回 true
}

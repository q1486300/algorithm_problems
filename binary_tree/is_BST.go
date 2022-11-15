package binary_tree

import "math"

func IsBST1(head *Node) bool {
	if head == nil {
		return true
	}
	var arr []*Node
	in(head, &arr)
	for i := 1; i < len(arr); i++ {
		if arr[i].value <= arr[i-1].value {
			return false
		}
	}
	return true
}

func in(head *Node, arr *[]*Node) {
	if head == nil {
		return
	}
	in(head.left, arr)
	*arr = append(*arr, head)
	in(head.right, arr)
}

func IsBST2(head *Node) bool {
	if head == nil {
		return true
	}
	return processIsBST(head).isBST
}

func processIsBST(head *Node) *IsBSTInfo {
	if head == nil {
		return nil
	}
	leftInfo := processIsBST(head.left)
	rightInfo := processIsBST(head.right)

	max := head.value
	if leftInfo != nil {
		max = int(math.Max(float64(max), float64(leftInfo.max)))
	}
	if rightInfo != nil {
		max = int(math.Max(float64(max), float64(rightInfo.max)))
	}

	min := head.value
	if leftInfo != nil {
		min = int(math.Min(float64(min), float64(leftInfo.min)))
	}
	if rightInfo != nil {
		min = int(math.Min(float64(min), float64(rightInfo.min)))
	}

	isBST := true
	if leftInfo != nil && !leftInfo.isBST {
		isBST = false
	}
	if rightInfo != nil && !rightInfo.isBST {
		isBST = false
	}
	if leftInfo != nil && leftInfo.max >= head.value {
		isBST = false
	}
	if rightInfo != nil && rightInfo.min <= head.value {
		isBST = false
	}
	return NewIsBSTINFO(isBST, max, min)
}

type IsBSTInfo struct {
	isBST    bool
	max, min int
}

func NewIsBSTINFO(i bool, max, min int) *IsBSTInfo {
	return &IsBSTInfo{
		isBST: i,
		max:   max,
		min:   min,
	}
}

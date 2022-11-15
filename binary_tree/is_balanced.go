package binary_tree

import "math"

func IsBalanced1(head *Node) bool {
	ans := make([]bool, 1)
	ans[0] = true
	processIsBalanced1(head, ans)
	return ans[0]
}

func processIsBalanced1(head *Node, ans []bool) int {
	if !ans[0] || head == nil {
		return -1
	}
	leftHeight := processIsBalanced1(head.left, ans)
	rightHeight := processIsBalanced1(head.right, ans)

	if math.Abs(float64(leftHeight-rightHeight)) > 1 {
		ans[0] = false
	}
	return int(math.Max(float64(leftHeight), float64(rightHeight)) + 1)
}

func IsBalanced2(head *Node) bool {
	return processIsBalanced2(head).isBalanced
}

func processIsBalanced2(head *Node) IsBalancedInfo {
	if head == nil {
		return NewIsBalancedInfo(true, 0)
	}
	leftInfo := processIsBalanced2(head.left)
	rightInfo := processIsBalanced2(head.right)

	height := int(math.Max(float64(leftInfo.height), float64(rightInfo.height))) + 1
	isBalanced := true
	if !leftInfo.isBalanced {
		isBalanced = false
	}
	if !rightInfo.isBalanced {
		isBalanced = false
	}
	if math.Abs(float64(leftInfo.height-rightInfo.height)) > 1 {
		isBalanced = false
	}
	return NewIsBalancedInfo(isBalanced, height)
}

type IsBalancedInfo struct {
	isBalanced bool
	height     int
}

func NewIsBalancedInfo(isBalanced bool, height int) IsBalancedInfo {
	return IsBalancedInfo{
		isBalanced: isBalanced,
		height:     height,
	}
}

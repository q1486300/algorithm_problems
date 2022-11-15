package binary_tree

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

type Node struct {
	value       int
	left, right *Node
}

func NewNode(value int) *Node {
	return &Node{value: value}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// 以下為隨機產生二元樹
func GenerateRandomBST(maxLevel, maxValue int) *Node {
	return generate(1, maxLevel, maxValue)
}

func generate(level, maxLevel, maxValue int) *Node {
	if level > maxLevel || rand.Float32() < 0.5 {
		return nil
	}
	head := NewNode(rand.Intn(maxValue + 1))
	head.left = generate(level+1, maxLevel, maxValue)
	head.right = generate(level+1, maxLevel, maxValue)
	return head
}

type NodeWithParent struct {
	value               int
	left, right, parent *NodeWithParent
}

func NewNodeWithParent(value int) *NodeWithParent {
	return &NodeWithParent{value: value}
}

func GenerateRandomBSTWithParent(maxLevel, maxValue int) *NodeWithParent {
	return generate2(1, maxLevel, maxValue)
}

func generate2(level, maxLevel, maxValue int) *NodeWithParent {
	if level > maxLevel || rand.Float32() < 0.5 {
		return nil
	}
	head := NewNodeWithParent(rand.Intn(maxValue + 1))
	head.left = generate2(level+1, maxLevel, maxValue)
	if head.left != nil {
		head.left.parent = head
	}
	head.right = generate2(level+1, maxLevel, maxValue)
	if head.right != nil {
		head.right.parent = head
	}
	return head
}

func PickRandomOneInListAndTheNext(head *NodeWithParent) (node1 *NodeWithParent, node2 *NodeWithParent) {
	if head == nil {
		return nil, nil
	}
	var arr []*NodeWithParent
	generateAllNodesInOrderList(head, &arr)
	randomIndex := rand.Intn(len(arr))
	if randomIndex == len(arr)-1 {
		return arr[randomIndex], nil
	}
	return arr[randomIndex], arr[randomIndex+1]
}

func generateAllNodesInOrderList(head *NodeWithParent, arr *[]*NodeWithParent) {
	if head == nil {
		return
	}
	generateAllNodesInOrderList(head.left, arr)
	*arr = append(*arr, head)
	generateAllNodesInOrderList(head.right, arr)
}

// for test
func PickRandomOne(head *Node) *Node {
	if head == nil {
		return nil
	}
	var arr []*Node
	fillPrelist(head, &arr)
	randomIndex := rand.Intn(len(arr))
	return arr[randomIndex]
}

func fillPrelist(head *Node, arr *[]*Node) {
	if head == nil {
		return
	}
	*arr = append(*arr, head)
	fillPrelist(head.left, arr)
	fillPrelist(head.right, arr)
}

// 以下為印出二元樹的結構
func PrintTree(head *Node) {
	fmt.Println("Binary Tree:")
	printInOrder(head, 0, "H", 17)
	fmt.Println()
}

func printInOrder(head *Node, height int, to string, length int) {
	if head == nil {
		return
	}
	printInOrder(head.right, height+1, "v", length)

	val := fmt.Sprintf("%v %d %v", to, head.value, to)
	lenM := len(val)
	lenL := (length - lenM) / 2
	lenR := length - lenM - lenL
	val = fmt.Sprintf("%v %v %v", getSpace(lenL), val, getSpace(lenR))
	fmt.Println(getSpace(height*length) + val)

	printInOrder(head.left, height+1, "^", length)
}

func getSpace(num int) string {
	space := strings.Builder{}
	for i := 0; i < num; i++ {
		space.WriteString(" ")
	}
	return space.String()
}

func testPrintTree() {
	h := NewNode(1)
	h.left = NewNode(-222222222)
	h.right = NewNode(3)
	h.left.left = NewNode(math.MinInt)
	h.right.left = NewNode(55555555)
	h.right.right = NewNode(66)
	h.left.left.right = NewNode(777)
	PrintTree(h)

	h = NewNode(1)
	h.left = NewNode(2)
	h.right = NewNode(3)
	h.left.left = NewNode(4)
	h.right.left = NewNode(5)
	h.right.right = NewNode(6)
	h.left.left.right = NewNode(7)
	PrintTree(h)

	h = NewNode(1)
	h.left = NewNode(1)
	h.right = NewNode(1)
	h.left.left = NewNode(1)
	h.right.left = NewNode(1)
	h.right.right = NewNode(1)
	h.left.left.right = NewNode(1)
	PrintTree(h)
}

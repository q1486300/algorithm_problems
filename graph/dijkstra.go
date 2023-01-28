package graph

import (
	"container/heap"
	"math"
)

func Dijkstra1(from *Node) map[*Node]int {
	// 從 from 節點到 Key 節點的最短距離是 Value
	distanceMap := make(map[*Node]int)
	distanceMap[from] = 0
	// 已經選過(處理過)的節點
	selectedNodes := make(map[*Node]struct{})
	minNode := getMinDistanceAndUnselectedNode(distanceMap, selectedNodes)
	for minNode != nil {
		// 原始點 -> minNode(跳轉點) 最小距離 distance
		distance := distanceMap[minNode]
		for _, edge := range minNode.edges {
			toNode := edge.to
			preMin, ok := distanceMap[toNode]
			if !ok {
				distanceMap[toNode] = distance + edge.weight
			} else {
				distanceMap[toNode] = getMin(preMin, distance+edge.weight)
			}
		}
		selectedNodes[minNode] = struct{}{}
		minNode = getMinDistanceAndUnselectedNode(distanceMap, selectedNodes)
	}
	return distanceMap
}

func getMinDistanceAndUnselectedNode(distanceMap map[*Node]int, selectedNodes map[*Node]struct{}) *Node {
	var minNode *Node
	minDistance := math.MaxInt
	for node, distance := range distanceMap {
		_, ok := selectedNodes[node]
		if !ok && distance < minDistance {
			minNode = node
			minDistance = distance
		}
	}
	return minNode
}

func Dijkstra2(head *Node) map[*Node]int {
	nodeDistanceHeap := NewNodeDistanceHeap()
	heap.Push(nodeDistanceHeap, NewNodeDistance(head, 0))

	result := make(map[*Node]int)
	for nodeDistanceHeap.Len() != 0 {
		nodeDistance := heap.Pop(nodeDistanceHeap).(NodeDistance)
		curNode := nodeDistance.node
		distance := nodeDistance.distance
		for _, edge := range curNode.edges {
			heap.Push(nodeDistanceHeap, NewNodeDistance(edge.to, distance+edge.weight))
		}
		result[curNode] = distance
	}
	return result
}

type NodeDistance struct {
	node     *Node
	distance int
}

type NodeDistanceHeap struct {
	nodeDistances []NodeDistance
	indexMap      map[*Node]int
}

func NewNodeDistance(node *Node, distance int) NodeDistance {
	return NodeDistance{
		node:     node,
		distance: distance,
	}
}

func NewNodeDistanceHeap() *NodeDistanceHeap {
	pq := &NodeDistanceHeap{
		nodeDistances: make([]NodeDistance, 1),
		indexMap:      make(map[*Node]int),
	}
	heap.Init(pq)
	return pq
}

func (n NodeDistanceHeap) Len() int {
	return len(n.nodeDistances)
}

func (n NodeDistanceHeap) Less(i, j int) bool {
	return n.nodeDistances[i].distance < n.nodeDistances[j].distance
}

func (n NodeDistanceHeap) Swap(i, j int) {
	n.indexMap[n.nodeDistances[i].node] = j
	n.indexMap[n.nodeDistances[j].node] = i
	n.nodeDistances[i], n.nodeDistances[j] = n.nodeDistances[j], n.nodeDistances[i]
}

func (n *NodeDistanceHeap) Push(x any) {
	nodeDistance := x.(NodeDistance)
	index, ok := n.indexMap[nodeDistance.node]
	if !ok { // node 還沒進到 heap 過
		index = len(n.nodeDistances)
		n.nodeDistances = append(n.nodeDistances, nodeDistance)
		n.indexMap[nodeDistance.node] = index
		return
	}
	if index != -1 { // node 已經進到 heap 裡，並且還沒被選擇過
		preNodeDistance := n.nodeDistances[index]
		preNodeDistance.distance = getMin(preNodeDistance.distance, nodeDistance.distance)
		heap.Fix(n, index)
	}
}

func (n *NodeDistanceHeap) Pop() any {
	length := len(n.nodeDistances)
	nodeDistance := n.nodeDistances[length-1]

	n.nodeDistances = n.nodeDistances[:length-1]
	n.indexMap[nodeDistance.node] = -1

	return nodeDistance
}

func getMin(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

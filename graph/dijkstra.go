package graph

import (
	"container/heap"
	"math"
)

func Dijkstra1(from *Node) map[*Node]int {
	distanceMap := make(map[*Node]int)
	distanceMap[from] = 0
	// 已經選過(處理過)的節點
	selectedNodes := make(map[*Node]bool)
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
				distanceMap[toNode] = int(math.Min(float64(preMin), float64(distance+edge.weight)))
			}
		}
		selectedNodes[minNode] = true
		minNode = getMinDistanceAndUnselectedNode(distanceMap, selectedNodes)
	}
	return distanceMap
}

func getMinDistanceAndUnselectedNode(distanceMap map[*Node]int, touchedNodes map[*Node]bool) *Node {
	var minNode *Node
	minDistance := math.MaxInt
	for node, distance := range distanceMap {
		_, ok := touchedNodes[node]
		if !ok && distance < minDistance {
			minNode = node
			minDistance = distance
		}
	}
	return minNode
}

func Dijkstra2(head *Node) map[*Node]int {
	nodeHeap := NewNodeHeap()
	heap.Push(nodeHeap, NewNodeRecord(head, 0))
	result := make(map[*Node]int)
	for nodeHeap.Len() != 0 {
		record := heap.Pop(nodeHeap).(NodeRecord)
		cur := record.node
		distance := record.distance
		for _, edge := range cur.edges {
			heap.Push(nodeHeap, NewNodeRecord(edge.to, edge.weight+distance))
		}
		result[cur] = distance
	}
	return result
}

type NodeRecord struct {
	node     *Node
	distance int
	index    int
}

type NodeHeap struct {
	nodeRecords []NodeRecord
	indexMap    map[*Node]int
}

func NewNodeRecord(node *Node, distance int) NodeRecord {
	return NodeRecord{
		node:     node,
		distance: distance,
	}
}

func NewNodeHeap() *NodeHeap {
	nodeHeap := &NodeHeap{
		nodeRecords: make([]NodeRecord, 0),
		indexMap:    make(map[*Node]int),
	}
	heap.Init(nodeHeap)
	return nodeHeap
}

func (n NodeHeap) Len() int {
	return len(n.nodeRecords)
}

func (n NodeHeap) Less(i, j int) bool {
	return n.nodeRecords[i].distance < n.nodeRecords[j].distance
}

func (n NodeHeap) Swap(i, j int) {
	n.nodeRecords[i], n.nodeRecords[j] = n.nodeRecords[j], n.nodeRecords[i]
	n.nodeRecords[i].index = j
	n.nodeRecords[j].index = i
	n.indexMap[n.nodeRecords[i].node] = j
	n.indexMap[n.nodeRecords[j].node] = i
}

func (n *NodeHeap) Push(x any) {
	nodeRecord := x.(NodeRecord)
	index, ok := n.indexMap[nodeRecord.node]
	if !ok {
		nodeRecord.index = len(n.nodeRecords)
		n.nodeRecords = append(n.nodeRecords, nodeRecord)
		n.indexMap[nodeRecord.node] = nodeRecord.index
		return
	}
	if index != -1 {
		node := n.nodeRecords[index]
		node.distance = int(math.Min(float64(n.nodeRecords[index].distance), float64(nodeRecord.distance)))
		heap.Fix(n, index)
	}
}

func (n *NodeHeap) Pop() any {
	length := len(n.nodeRecords)
	nodeRecord := n.nodeRecords[length-1]

	n.nodeRecords = n.nodeRecords[:length-1]

	n.indexMap[nodeRecord.node] = -1
	return nodeRecord
}

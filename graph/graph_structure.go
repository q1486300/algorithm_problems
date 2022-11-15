package graph

import "container/heap"

type Graph struct {
	nodes map[int]*Node
	edges map[*Edge]bool
}

func (g Graph) GetAllNodes() []*Node {
	var nodes []*Node
	for _, node := range g.nodes {
		nodes = append(nodes, node)
	}
	return nodes
}

type Node struct {
	value   int
	in, out int
	nexts   []*Node
	edges   []*Edge
}

type Edge struct {
	weight   int
	from, to *Node
}

type EdgePriorityQueue []*Edge

func (e EdgePriorityQueue) Len() int {
	return len(e)
}

func (e EdgePriorityQueue) Less(i, j int) bool {
	return e[i].weight < e[j].weight
}

func (e EdgePriorityQueue) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (e *EdgePriorityQueue) Push(x any) {
	*e = append(*e, x.(*Edge))
}

func (e *EdgePriorityQueue) Pop() any {
	old := *e
	tmp := old[len(*e)-1]
	*e = old[:len(*e)-1]
	return tmp
}

func NewGraph() Graph {
	return Graph{
		nodes: make(map[int]*Node),
		edges: make(map[*Edge]bool),
	}
}

func NewNode(value int) *Node {
	return &Node{value: value}
}

func NewEdge(weight int, from, to *Node) *Edge {
	return &Edge{
		weight: weight,
		from:   from,
		to:     to,
	}
}

func NewEdgePriorityQueue() *EdgePriorityQueue {
	pq := &EdgePriorityQueue{}
	heap.Init(pq)
	return pq
}

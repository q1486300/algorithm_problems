package graph

import (
	"container/heap"
	"container/list"
)

// undirected graph only

// Union-Find Set
type UnionFind struct {
	// key: 某一個節點，value: key節點往上的節點
	fatherMap map[*Node]*Node
	// key: 某一個集合的代表節點，value: key所在集合的節點個數
	sizeMap map[*Node]int
}

func NewUnionFind() UnionFind {
	return UnionFind{
		fatherMap: make(map[*Node]*Node),
		sizeMap:   make(map[*Node]int),
	}
}

func (u UnionFind) MakeSets(nodes []*Node) {
	u.fatherMap = make(map[*Node]*Node)
	u.sizeMap = make(map[*Node]int)
	for _, node := range nodes {
		u.fatherMap[node] = node
		u.sizeMap[node] = 1
	}
}

func (u UnionFind) findFather(cur *Node) *Node {
	path := list.New()
	for cur != u.fatherMap[cur] {
		path.PushBack(cur)
		cur = u.fatherMap[cur]
	}
	for path.Len() != 0 {
		u.fatherMap[path.Back().Value.(*Node)] = cur
		path.Remove(path.Back())
	}
	return cur
}

func (u UnionFind) IsSameSet(a, b *Node) bool {
	return u.findFather(a) == u.findFather(b)
}

func (u UnionFind) Union(a, b *Node) {
	if a == nil || b == nil {
		return
	}
	aF := u.findFather(a)
	bF := u.findFather(b)
	if aF != bF {
		aSetSize := u.sizeMap[aF]
		bSetSize := u.sizeMap[bF]
		if aSetSize <= bSetSize {
			u.fatherMap[aF] = bF
			u.sizeMap[bF] = aSetSize + bSetSize
			delete(u.sizeMap, aF)
		} else {
			u.fatherMap[bF] = aF
			u.sizeMap[aF] = aSetSize + bSetSize
			delete(u.sizeMap, bF)
		}
	}
}

func (u UnionFind) Sets() int {
	return len(u.sizeMap)
}

func KruskalMST(graph Graph) map[*Edge]struct{} {
	unionFind := NewUnionFind()
	unionFind.MakeSets(graph.GetAllNodes())
	// 從小的邊到大的邊，一次彈出，最小堆積樹
	priorityQueue := NewEdgePriorityQueue()
	for edge := range graph.edges { // M 條邊
		heap.Push(priorityQueue, edge) // O(logM)
	}
	result := make(map[*Edge]struct{})
	for priorityQueue.Len() != 0 { // M 條邊
		edge := heap.Pop(priorityQueue).(*Edge)
		if !unionFind.IsSameSet(edge.from, edge.to) { // O(1)
			result[edge] = struct{}{}
			unionFind.Union(edge.from, edge.to)
		}
	}
	return result
}

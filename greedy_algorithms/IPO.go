package greedy_algorithms

import "container/heap"

// 測試連結: https://leetcode.com/problems/ipo/
// 最多 k 個項目
// w 是初始資金
// profits、capital 長度相等
// 返回最終最大的資金
func findMaximizedCapital(k, w int, profits, capital []int) int {
	minCostPQ := NewMinCostPriorityQueue()
	maxProfitPQ := NewMaxProfitPriorityQueue()
	for i := 0; i < len(profits); i++ {
		heap.Push(minCostPQ, NewProgramIPO(profits[i], capital[i]))
	}
	for i := 0; i < k; i++ {
		for minCostPQ.Len() != 0 && (*minCostPQ)[0].c <= w {
			heap.Push(maxProfitPQ, heap.Pop(minCostPQ).(ProgramIPO))
		}
		if maxProfitPQ.Len() == 0 {
			return w
		}
		w += heap.Pop(maxProfitPQ).(ProgramIPO).p
	}
	return w
}

type ProgramIPO struct {
	p, c int
}

func NewProgramIPO(p, c int) ProgramIPO {
	return ProgramIPO{
		p: p,
		c: c,
	}
}

type MinCostProgramIPO []ProgramIPO

func NewMinCostPriorityQueue() *MinCostProgramIPO {
	pQ := &MinCostProgramIPO{}
	heap.Init(pQ)
	return pQ
}

func (m MinCostProgramIPO) Len() int {
	return len(m)
}

func (m MinCostProgramIPO) Less(i, j int) bool {
	return m[i].c < m[j].c
}

func (m MinCostProgramIPO) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MinCostProgramIPO) Push(x any) {
	*m = append(*m, x.(ProgramIPO))
}

func (m *MinCostProgramIPO) Pop() any {
	n := len(*m)
	result := (*m)[n-1]
	*m = (*m)[:n-1]
	return result
}

type MaxProfitProgramIPO []ProgramIPO

func NewMaxProfitPriorityQueue() *MaxProfitProgramIPO {
	pQ := &MaxProfitProgramIPO{}
	heap.Init(pQ)
	return pQ
}

func (m MaxProfitProgramIPO) Len() int {
	return len(m)
}

func (m MaxProfitProgramIPO) Less(i, j int) bool {
	return m[i].p > m[j].p
}

func (m MaxProfitProgramIPO) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MaxProfitProgramIPO) Push(x any) {
	*m = append(*m, x.(ProgramIPO))
}

func (m *MaxProfitProgramIPO) Pop() any {
	n := len(*m)
	result := (*m)[n-1]
	*m = (*m)[:n-1]
	return result
}

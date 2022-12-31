package greedy_algorithms

import (
	"container/heap"
	"math"
)

func LessMoney1(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return 0
	}
	return processLessMoney(arr, 0)
}

// 等待合併的數都在 arr 裡，pre 之前的合併行為產生了多少總代價
// arr 中只剩一個數字的時候，停止合併，返回最小的總代價
func processLessMoney(arr []int, pre int) int {
	if len(arr) == 1 {
		return pre
	}
	ans := math.MaxInt
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			ans = getMin(ans, processLessMoney(copyAndMergeTwo(arr, i, j), pre+arr[i]+arr[j]))
		}
	}
	return ans
}

func copyAndMergeTwo(arr []int, i int, j int) []int {
	result := make([]int, len(arr)-1)
	index := 0
	for k, cur := range arr {
		if k != i && k != j {
			result[index] = cur
			index++
		}
	}
	result[index] = arr[i] + arr[j]
	return result
}

func getMin(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func LessMoney2(arr []int) int {
	pQ := NewMoneyPriorityQueue()
	for _, cur := range arr {
		heap.Push(pQ, cur)
	}
	sum, cur := 0, 0
	for pQ.Len() > 1 {
		cur = heap.Pop(pQ).(int) + heap.Pop(pQ).(int)
		sum += cur
		heap.Push(pQ, cur)
	}
	return sum
}

type MoneyPriorityQueue []int

func NewMoneyPriorityQueue() *MoneyPriorityQueue {
	pQ := &MoneyPriorityQueue{}
	heap.Init(pQ)
	return pQ
}

func (m MoneyPriorityQueue) Len() int {
	return len(m)
}

func (m MoneyPriorityQueue) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m MoneyPriorityQueue) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MoneyPriorityQueue) Push(x any) {
	*m = append(*m, x.(int))
}

func (m *MoneyPriorityQueue) Pop() any {
	n := len(*m)
	result := (*m)[n-1]
	*m = (*m)[:n-1]
	return result
}

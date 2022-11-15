package brute_force_recursion

import (
	"math"
)

// 所有貨的重量和價值，都在 w 和 v 切片裡
// 其中沒有負數
// bag 背包容量，不能超過這個載重
// 返回: 不超重的情況下，能夠得到的最大價值
func MaxValue(w, v []int, bag int) int {
	if w == nil || v == nil || len(w) != len(v) || len(w) == 0 {
		return 0
	}
	return processMaxValue(w, v, 0, bag)
}

// index 0 ~ N
// rest 負數 ~ bag
func processMaxValue(w, v []int, index, rest int) int {
	if rest < 0 {
		return -1
	}
	if index == len(w) {
		return 0
	}
	p1 := processMaxValue(w, v, index+1, rest)
	p2 := 0
	next := processMaxValue(w, v, index+1, rest-w[index])
	if next != -1 {
		p2 = v[index] + next
	}
	return int(math.Max(float64(p1), float64(p2)))
}

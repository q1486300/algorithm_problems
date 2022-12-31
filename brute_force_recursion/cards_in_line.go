package brute_force_recursion

// 根據規則，返回獲勝者的分數
func Win(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return 0
	}
	first := f(arr, 0, len(arr)-1)
	second := g(arr, 0, len(arr)-1)
	return getMax(first, second)
}

// arr[L..R]，先手獲得的最好分數返回
func f(arr []int, L, R int) int {
	if L == R {
		return arr[L]
	}
	p1 := arr[L] + g(arr, L+1, R)
	p2 := arr[R] + g(arr, L, R-1)
	return getMax(p1, p2)
}

// arr[L..R]，後手獲得的最好分數返回
func g(arr []int, L, R int) int {
	if L == R {
		return 0
	}
	p1 := f(arr, L+1, R) // 對手拿走了 L 位置上的數
	p2 := f(arr, L, R-1) // 對手拿走了 R 位置上的數
	return getMin(p1, p2)
}

func getMax(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func getMin(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

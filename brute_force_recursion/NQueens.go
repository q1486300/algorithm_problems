package brute_force_recursion

import "math"

func NQueens1(n int) int {
	if n < 1 {
		return 0
	}
	record := make([]int, n)
	return processNQueens1(0, record, n)
}

// 目前來到第 i 列，一共有 0 ~ n-1 列
// 在 i 列上放皇后，所有行都嘗試
// 必須保證跟之前所有的皇后不衝突(不同列、不同行、不同斜線)
// record[x] = y，表示之前第 x 列的皇后，放在了 y 行上
// 返回: 不關心 i 以上發生了什麼，i.. 後續有多少合法的方法數
func processNQueens1(i int, record []int, n int) int {
	if i == n {
		return 1
	}
	res := 0
	// i 列的皇后，放在 j 行
	for j := 0; j < n; j++ {
		if isValid(record, i, j) {
			record[i] = j
			res += processNQueens1(i+1, record, n)
		}
	}
	return res
}

func isValid(record []int, i, j int) bool {
	// 0..i-1
	for k := 0; k < i; k++ {
		// 判斷是否同行 || 同斜線
		if record[k] == j || math.Abs(float64(record[k]-j)) == math.Abs(float64(k-i)) {
			return false
		}
	}
	return true
}

// 請不要超過 32 皇后問題
func NQueens2(n int) int {
	if n < 1 || n > 32 {
		return 0
	}
	// 如果是 7 皇后問題，limit 最右(位) 7 個 1，其他都是 0
	var limit int
	if n == 32 {
		limit = -1
	} else {
		limit = (1 << n) - 1
	}

	return processNQueens2(limit, 0, 0, 0)
}

// 7 皇后問題
// limit: 0...0 1 1 1 1 1 1 1
// 之前皇后的行限制: colLim
// 之前皇后的左下對角線限制: leftDiaLim
// 之前皇后的右下對角線限制: rightDiaLim
func processNQueens2(limit, colLim, leftDiaLim, rightDiaLim int) int {
	if colLim == limit {
		return 1
	}
	// pos 中所有是 1 的位置，表示可以嘗試皇后的位置
	pos := limit & (^(colLim | leftDiaLim | rightDiaLim))
	mostRightOne := 0
	res := 0
	for pos != 0 {
		mostRightOne = pos & (^pos + 1)
		pos = pos - mostRightOne
		res += processNQueens2(limit, colLim|mostRightOne, (leftDiaLim|mostRightOne)<<1,
			(rightDiaLim|mostRightOne)>>1)
	}
	return res
}

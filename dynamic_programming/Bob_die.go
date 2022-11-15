package dynamic_programming

import "math"

func LivePossibility1(row, col, k, N, M int) float64 {
	return float64(processLP1(row, col, k, N, M)) / math.Pow(4, float64(k))
}

// 目前在 (row, col) 位置，還有 rest 步需要走
// 如果走完了還在棋盤中就獲得 1 個生存點，返回總共的生存點
func processLP1(row int, col int, rest int, N int, M int) int64 {
	if row < 0 || row == N || col < 0 || col == M {
		return 0
	}
	// 還在棋盤中
	if rest == 0 {
		return 1
	}
	// 還在棋盤中，還有步數要走
	up := processLP1(row-1, col, rest-1, N, M)
	down := processLP1(row+1, col, rest-1, N, M)
	left := processLP1(row, col-1, rest-1, N, M)
	right := processLP1(row, col+1, rest-1, N, M)
	return up + down + left + right
}

func LivePossibility2(row, col, k, N, M int) float64 {
	dp := make([][][]int64, N)
	for i := range dp {
		dp[i] = make([][]int64, M)
		for j := range dp[i] {
			dp[i][j] = make([]int64, k+1)
			dp[i][j][0] = 1
		}
	}
	for step := 1; step <= k; step++ {
		for r := 0; r < N; r++ {
			for c := 0; c < M; c++ {
				dp[r][c][step] = getValue(dp, N, M, r-1, c, step-1)
				dp[r][c][step] += getValue(dp, N, M, r+1, c, step-1)
				dp[r][c][step] += getValue(dp, N, M, r, c-1, step-1)
				dp[r][c][step] += getValue(dp, N, M, r, c+1, step-1)
			}
		}
	}
	return float64(dp[row][col][k]) / math.Pow(4, float64(k))
}

func getValue(dp [][][]int64, N, M, r, c, step int) int64 {
	if r < 0 || r == N || c < 0 || c == M {
		return 0
	}
	return dp[r][c][step]
}

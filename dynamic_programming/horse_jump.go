package dynamic_programming

// 目前來到的位置是 (x, y)
// 還剩下 rest 步需要跳
// 跳完 rest 步，剛好跳到 (a, b) 的方法數是多少
// 10 * 9
func HorseJump(a, b, k int) int {
	return processHorseJump(0, 0, k, a, b)
}

func processHorseJump(row, col, rest, a, b int) int {
	if row < 0 || row > 9 || col < 0 || col > 8 {
		return 0
	}
	if rest == 0 {
		if row == a && col == b {
			return 1
		} else {
			return 0
		}
	}
	ways := processHorseJump(row+2, col+1, rest-1, a, b)
	ways += processHorseJump(row+1, col+2, rest-1, a, b)
	ways += processHorseJump(row-1, col+2, rest-1, a, b)
	ways += processHorseJump(row-2, col+1, rest-1, a, b)
	ways += processHorseJump(row-2, col-1, rest-1, a, b)
	ways += processHorseJump(row-1, col-2, rest-1, a, b)
	ways += processHorseJump(row+1, col-2, rest-1, a, b)
	ways += processHorseJump(row+2, col-1, rest-1, a, b)
	return ways
}

func HorseJumpDP(a, b, k int) int {
	dp := make([][][]int, 10)
	for i := range dp {
		dp[i] = make([][]int, 9)
		for j := range dp[i] {
			dp[i][j] = make([]int, k+1)
		}
	}
	dp[a][b][0] = 1
	for step := 1; step <= k; step++ {
		for row := 0; row < 10; row++ {
			for col := 0; col < 9; col++ {
				ways := pick(dp, row+2, col+1, step-1)
				ways += pick(dp, row+1, col+2, step-1)
				ways += pick(dp, row-1, col+2, step-1)
				ways += pick(dp, row-2, col+1, step-1)
				ways += pick(dp, row-2, col-1, step-1)
				ways += pick(dp, row-1, col-2, step-1)
				ways += pick(dp, row+1, col-2, step-1)
				ways += pick(dp, row+2, col-1, step-1)
				dp[row][col][step] = ways
			}
		}
	}
	return dp[0][0][k]
}

func pick(dp [][][]int, row, col, step int) int {
	if row < 0 || row > 9 || col < 0 || col > 8 {
		return 0
	}
	return dp[row][col][step]
}

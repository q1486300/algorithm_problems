package dynamic_programming

// 目前來到的位置是 (x, y)
// 還剩下 rest 步需要跳
// 跳完 rest 步，剛好跳到 (a, b) 的方法數是多少
// 10 * 9
func HorseJump(a, b, k int) int {
	return processHorseJump(0, 0, k, a, b)
}

func processHorseJump(x, y, rest, a, b int) int {
	if x < 0 || x > 9 || y < 0 || y > 8 {
		return 0
	}
	if rest == 0 {
		if x == a && y == b {
			return 1
		} else {
			return 0
		}
	}
	ways := processHorseJump(x+2, y+1, rest-1, a, b)
	ways += processHorseJump(x+1, y+2, rest-1, a, b)
	ways += processHorseJump(x-1, y+2, rest-1, a, b)
	ways += processHorseJump(x-2, y+1, rest-1, a, b)
	ways += processHorseJump(x-2, y-1, rest-1, a, b)
	ways += processHorseJump(x-1, y-2, rest-1, a, b)
	ways += processHorseJump(x+1, y-2, rest-1, a, b)
	ways += processHorseJump(x+2, y-1, rest-1, a, b)
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
		for x := 0; x < 10; x++ {
			for y := 0; y < 9; y++ {
				ways := pick(dp, x+2, y+1, step-1)
				ways += pick(dp, x+1, y+2, step-1)
				ways += pick(dp, x-1, y+2, step-1)
				ways += pick(dp, x-2, y+1, step-1)
				ways += pick(dp, x-2, y-1, step-1)
				ways += pick(dp, x-1, y-2, step-1)
				ways += pick(dp, x+1, y-2, step-1)
				ways += pick(dp, x+2, y-1, step-1)
				dp[x][y][step] = ways
			}
		}
	}
	return dp[0][0][k]
}

func pick(dp [][][]int, x, y, step int) int {
	if x < 0 || x > 9 || y < 0 || y > 8 {
		return 0
	}
	return dp[x][y][step]
}

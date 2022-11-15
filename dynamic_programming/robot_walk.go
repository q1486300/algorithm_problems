package dynamic_programming

func RobotWalk1(N, start, aim, K int) int {
	if N < 2 || start < 1 || start > N || aim < 1 || aim > N || K < 1 {
		return -1
	}
	return processRobotWalk1(start, K, aim, N)
}

// 機器人目前來到的位置是 cur
// 機器人還有 rest 步需要走
// 最終的目標是 aim
// 總共有 1 ~ N 位置
// 返回: 機器人從 cur 出發，走過 rest 步之後，最終停在 aim 的方法數
func processRobotWalk1(cur, rest, aim, N int) int {
	if rest == 0 { // 如果已經不需要走了，亦即走完了
		if cur == aim {
			return 1
		} else {
			return 0
		}
	}
	if cur == 1 { // 1 -> 2
		return processRobotWalk1(2, rest-1, aim, N)
	}
	if cur == N { // N-1 <- N
		return processRobotWalk1(N-1, rest-1, aim, N)
	}
	// 中間位置
	return processRobotWalk1(cur-1, rest-1, aim, N) + processRobotWalk1(cur+1, rest-1, aim, N)
}

func RobotWalk2(N, start, aim, K int) int {
	if N < 2 || start < 1 || start > N || aim < 1 || aim > N || K < 1 {
		return -1
	}
	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, K+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	// dp 就是緩存表
	// dp[cur][rest] == -1 -> processRobotWalk2(cur, rest) 之前沒算過
	// dp[cur][rest] != -1 -> processRobotWalk2(cur, rest) 之前算過，返回值: dp[cur][rest]
	// N+1 * K+1
	return processRobotWalk2(start, K, aim, N, dp)
}

// cur 範圍: 1 ~ N
// rest 範圍: 0 ~ K
func processRobotWalk2(cur, rest, aim, N int, dp [][]int) int {
	if dp[cur][rest] != -1 {
		return dp[cur][rest]
	}
	// 之前沒算過
	ans := 0
	if rest == 0 {
		if cur == aim {
			ans = 1
		} else {
			ans = 0
		}
	} else if cur == 1 {
		ans = processRobotWalk2(2, rest-1, aim, N, dp)
	} else if cur == N {
		ans = processRobotWalk2(N-1, rest-1, aim, N, dp)
	} else {
		ans = processRobotWalk2(cur-1, rest-1, aim, N, dp) + processRobotWalk2(cur+1, rest-1, aim, N, dp)
	}
	dp[cur][rest] = ans
	return ans
}

func RobotWalk3(N, start, aim, K int) int {
	if N < 2 || start < 1 || start > N || aim < 1 || aim > N || K < 1 {
		return -1
	}
	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, K+1)
	}
	dp[aim][0] = 1
	for rest := 1; rest <= K; rest++ {
		dp[1][rest] = dp[2][rest-1]
		for cur := 2; cur < N; cur++ {
			dp[cur][rest] = dp[cur-1][rest-1] + dp[cur+1][rest-1]
		}
		dp[N][rest] = dp[N-1][rest-1]
	}
	return dp[start][K]
}

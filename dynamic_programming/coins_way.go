package dynamic_programming

func CoinsWay(arr []int, aim int) int {
	if arr == nil || len(arr) == 0 || aim < 0 {
		return 0
	}
	return processCoinsWay(arr, 0, aim)
}

// arr[index..] 所有的面值，每一個面值都可以任意選擇張數，組成剛好 rest 這麼多錢，方法數是多少?
func processCoinsWay(arr []int, index, rest int) int {
	if index == len(arr) { // 沒有面值可選了
		if rest == 0 {
			return 1
		} else {
			return 0
		}
	}
	ways := 0
	for nums := 0; nums*arr[index] <= rest; nums++ {
		ways += processCoinsWay(arr, index+1, rest-(nums*arr[index]))
	}
	return ways
}

func CoinsWayDP1(arr []int, aim int) int {
	if arr == nil || len(arr) == 0 || aim < 0 {
		return 0
	}
	N := len(arr)
	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, aim+1)
	}
	dp[N][0] = 1
	for index := N - 1; index >= 0; index-- {
		for rest := 0; rest <= aim; rest++ {
			ways := 0
			for nums := 0; nums*arr[index] <= rest; nums++ {
				ways += dp[index+1][rest-(nums*arr[index])]
			}
			dp[index][rest] = ways
		}
	}
	return dp[0][aim]
}

func CoinsWayDP2(arr []int, aim int) int {
	if arr == nil || len(arr) == 0 || aim < 0 {
		return 0
	}
	N := len(arr)
	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, aim+1)
	}
	dp[N][0] = 1
	for index := N - 1; index >= 0; index-- {
		for rest := 0; rest <= aim; rest++ {
			dp[index][rest] = dp[index+1][rest]
			if rest-arr[index] >= 0 {
				dp[index][rest] += dp[index][rest-arr[index]]
			}
		}
	}
	return dp[0][aim]
}

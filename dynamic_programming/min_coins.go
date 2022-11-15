package dynamic_programming

import "math"

func MinCoins(arr []int, aim int) int {
	return processMinCoins1(arr, 0, aim)
}

// arr[index..] 面值，每種面值張數自由選擇
// 找出 rest 剛好這麼多錢，返回最小張數
// 用 math.MaxInt 標記: 怎樣都無法找出 rest 剛好這麼多錢
func processMinCoins1(arr []int, index, rest int) int {
	if index == len(arr) {
		if rest == 0 {
			return 0
		} else {
			return math.MaxInt
		}
	} else {
		ans := math.MaxInt
		for nums := 0; nums*arr[index] <= rest; nums++ {
			next := processMinCoins1(arr, index+1, rest-(nums*arr[index]))
			if next != math.MaxInt {
				ans = int(math.Min(float64(ans), float64(nums+next)))
			}
		}
		return ans
	}
}

func MinCoinsDP1(arr []int, aim int) int {
	if aim == 0 {
		return 0
	}
	N := len(arr)
	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, aim+1)
	}
	dp[N][0] = 0
	for j := 1; j <= aim; j++ {
		dp[N][j] = math.MaxInt
	}
	for index := N - 1; index >= 0; index-- {
		for rest := 0; rest <= aim; rest++ {
			ans := math.MaxInt
			for nums := 0; nums*arr[index] <= rest; nums++ {
				next := dp[index+1][rest-(nums*arr[index])]
				if next != math.MaxInt {
					ans = int(math.Min(float64(ans), float64(nums+next)))
				}
			}
			dp[index][rest] = ans
		}
	}
	return dp[0][aim]
}

func MinCoinsDP2(arr []int, aim int) int {
	if aim == 0 {
		return 0
	}
	N := len(arr)
	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, aim+1)
	}
	dp[N][0] = 0
	for j := 1; j <= aim; j++ {
		dp[N][j] = math.MaxInt
	}
	for index := N - 1; index >= 0; index-- {
		for rest := 0; rest <= aim; rest++ {
			dp[index][rest] = dp[index+1][rest]
			if rest-arr[index] >= 0 && dp[index][rest-arr[index]] != math.MaxInt {
				dp[index][rest] = int(math.Min(float64(dp[index][rest]), float64(dp[index][rest-arr[index]]+1)))
			}
		}
	}
	return dp[0][aim]
}

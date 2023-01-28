package brute_force_recursion

// str 只含有數字字元 0 ~ 9
// 返回多少種轉化方法
func Number(str string) int {
	if len(str) == 0 {
		return 0
	}
	intStr := make([]int, len(str))
	for i, cur := range str {
		intStr[i] = int(cur - 48) // '0' 的 ASCII碼 為 48
	}
	return processNumber(intStr, 0)
}

// str[0..i-1] 之前如何轉化已經做過決定了
// str[i..] 去轉化，返回有多少種轉化方法
func processNumber(str []int, i int) int {
	if i == len(str) {
		return 1
	}
	// i 沒到最後，說明後面還有字元
	if str[i] == 0 { // 之前的轉化，讓這次 i 位置上的 0 無法轉化成功
		return 0
	}
	// str[i] != '0'
	// 可能性一，i 單獨轉化
	ways := processNumber(str, i+1)
	// 可能性二，i 和 i + 1 位置上的數看成一組，並且小於 27
	if i+1 < len(str) && (str[i]*10)+(str[i+1]) < 27 {
		ways += processNumber(str, i+2)
	}
	return ways
}

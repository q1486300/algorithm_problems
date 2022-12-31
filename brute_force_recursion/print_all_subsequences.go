package brute_force_recursion

func PrintAllSubsequences(s string) []string {
	str := []rune(s)
	path := ""
	var ans []string
	processPAS(str, 0, &ans, path)
	return ans
}

// str 固定參數
// 來到了 str[index] 字元，index 是位置
// str[0..index-1] 已經走過了，之前的決定都在 path 上
// 之前的決定已經不能改變了，就是 path
// str[index..] 還能決定，之前已經確定，而後面還能自由選擇的話
// 把所有產生的子序列，放入 ans 中
func processPAS(str []rune, index int, ans *[]string, path string) {
	if index == len(str) {
		*ans = append(*ans, path)
		return
	}
	// 沒有要 index 位置上的字元
	no := path
	processPAS(str, index+1, ans, no)
	// 要 index 位置上的字元
	yes := path + string(str[index])
	processPAS(str, index+1, ans, yes)
}

func PrintAllSubsequencesNoRepeat(s string) []string {
	str := []rune(s)
	path := ""
	set := make(map[string]struct{})
	processPASNR(str, 0, set, path)
	var ans []string
	for cur, _ := range set {
		ans = append(ans, cur)
	}
	return ans
}

func processPASNR(str []rune, index int, set map[string]struct{}, path string) {
	if index == len(str) {
		set[path] = struct{}{}
		return
	}
	// 沒有要 index 位置上的字元
	no := path
	processPASNR(str, index+1, set, no)
	// 要 index 位置上的字元
	yes := path + string(str[index])
	processPASNR(str, index+1, set, yes)
}

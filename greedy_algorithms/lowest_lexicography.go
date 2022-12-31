package greedy_algorithms

import (
	"sort"
	"strings"
)

func LowestString1(strs []string) string {
	if strs == nil || len(strs) == 0 {
		return ""
	}
	ans := processLowestLexicography(strs)
	return ans[0]
}

func processLowestLexicography(strs []string) []string {
	var ans []string
	if len(strs) == 0 {
		ans = append(ans, "")
		return ans
	}
	for i, str := range strs {
		first := str
		nexts := removeIndexString(strs, i)
		next := processLowestLexicography(nexts)
		for _, cur := range next {
			ans = append(ans, first+cur)
		}
	}
	sort.Strings(ans)
	return ans
}

// {"abc", "cks", "bct"}
// 0 1 2
// removeIndexString(arr, 1) -> {"abc", "bct"}
func removeIndexString(strs []string, index int) []string {
	ans := make([]string, len(strs)-1)
	ansIndex := 0
	for i, str := range strs {
		if i != index {
			ans[ansIndex] = str
			ansIndex++
		}
	}
	return ans
}

type MyStrings []string

func (m MyStrings) Len() int {
	return len(m)
}

func (m MyStrings) Less(i, j int) bool {
	if m[i]+m[j] < m[j]+m[i] {
		return true
	}
	return false
}

func (m MyStrings) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func LowestString2(strs MyStrings) string {
	sort.Sort(strs)
	res := strings.Builder{}
	for _, str := range strs {
		res.WriteString(str)
	}
	return res.String()
}

package brute_force_recursion

func PrintAllPermutation1(s string) []string {
	var ans []string
	if len(s) == 0 {
		return ans
	}
	str := []rune(s)
	path := ""
	processPAP1(str, path, &ans)
	return ans
}

func processPAP1(rest []rune, path string, ans *[]string) {
	if len(rest) == 0 {
		*ans = append(*ans, path)
		return
	}
	for i := 0; i < len(rest); i++ {
		tmp := make([]rune, len(rest))
		copy(tmp, rest)
		cur := rest[i]
		rest = append(rest[:i], rest[i+1:]...)
		processPAP1(rest, path+string(cur), ans)
		rest = tmp
	}
}

func PrintAllPermutation2(s string) []string {
	var ans []string
	if len(s) == 0 {
		return ans
	}
	str := []rune(s)
	processPAP2(str, 0, &ans)
	return ans
}

func processPAP2(str []rune, index int, ans *[]string) {
	if index == len(str) {
		*ans = append(*ans, string(str))
		return
	}
	for i := index; i < len(str); i++ {
		swap(str, index, i)
		processPAP2(str, index+1, ans)
		swap(str, index, i)
	}
}

// PrintAllPermutation2 的去重版本
func PrintAllPermutation3(s string) []string {
	var ans []string
	if len(s) == 0 {
		return ans
	}
	str := []rune(s)
	processPAP3(str, 0, &ans)
	return ans
}

func processPAP3(str []rune, index int, ans *[]string) {
	if index == len(str) {
		*ans = append(*ans, string(str))
		return
	}
	visited := make([]bool, 256)
	for i := index; i < len(str); i++ {
		if !visited[str[i]] {
			visited[str[i]] = true
			swap(str, index, i)
			processPAP3(str, index+1, ans)
			swap(str, index, i)
		}
	}
}

func swap(str []rune, i, j int) {
	str[i], str[j] = str[j], str[i]
}

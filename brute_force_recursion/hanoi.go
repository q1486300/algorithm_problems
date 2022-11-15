package brute_force_recursion

import "fmt"

func Hanoi(n int) {
	if n > 0 {
		processHanoi(n, "left", "right", "mid")
	}
}

func processHanoi(n int, from, to, other string) {
	if n == 1 {
		fmt.Printf("Move 1 from %s to %s\n", from, to)
		return
	}
	processHanoi(n-1, from, other, to)
	fmt.Printf("Move %d from %s to %s\n", n, from, to)
	processHanoi(n-1, other, to, from)
}

package brute_force_recursion

import (
	"fmt"
	"testing"
)

func TestPrintAllSubsequences(t *testing.T) {
	test := "accc"
	ans1 := PrintAllSubsequences(test)
	ans2 := PrintAllSubsequencesNoRepeat(test)

	for _, str := range ans1 {
		fmt.Println(str)
	}
	fmt.Println("=====================")
	for _, str := range ans2 {
		fmt.Println(str)
	}
	fmt.Println("=====================")
}

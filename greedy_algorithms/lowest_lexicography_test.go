package greedy_algorithms

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestLowestString(t *testing.T) {
	testTimes := 50000
	arrLen := 6
	strLen := 5

	for i := 0; i < testTimes; i++ {
		arr1 := generateRandomStringArray(arrLen, strLen)
		arr2 := copyStringArray(arr1)

		s1 := LowestString1(arr1)
		s2 := LowestString2(arr2)

		if !assert.Equal(t, s1, s2) {
			for _, str := range arr1 {
				fmt.Printf("%s, ", str)
			}
			fmt.Println()
			break
		}
	}
}

func generateRandomString(strLen int) string {
	ans := make([]byte, rand.Intn(strLen)+1)
	for i := 0; i < len(ans); i++ {
		value := byte(rand.Intn(5))
		if rand.Float32() < 0.5 {
			ans[i] = 65 + value
		} else {
			ans[i] = 97 + value
		}
	}
	return string(ans)
}

func generateRandomStringArray(arrLen, strLen int) []string {
	ans := make([]string, rand.Intn(arrLen)+1)
	for i := 0; i < len(ans); i++ {
		ans[i] = generateRandomString(strLen)
	}
	return ans
}

func copyStringArray(arr []string) []string {
	ans := make([]string, len(arr))
	copy(ans, arr)
	return ans
}

package string

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"strings"
	"testing"
)

func TestKMP(t *testing.T) {
	testTimes := 5000000
	possibilities := 5
	strSize := 20
	matchSize := 5

	for i := 0; i < testTimes; i++ {
		str := getRandomString(possibilities, strSize)
		match := getRandomString(possibilities, matchSize)

		i1 := GetIndexOf(str, match)
		i2 := strings.Index(str, match)

		if !assert.Equal(t, i2, i1) {
			break
		}
	}
}

func getRandomString(possibilities, size int) string {
	ans := make([]byte, rand.Intn(size)+1)
	for i := 0; i < len(ans); i++ {
		ans[i] = byte(rand.Intn(possibilities)) + 'a'
	}
	return string(ans)
}

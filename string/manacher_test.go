package string

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestManacher(t *testing.T) {
	testTimes := 5000000
	possibilities := 5
	strSize := 20

	for i := 0; i < testTimes; i++ {
		str := getRandomString(possibilities, strSize)

		i1 := Manacher(str)
		i2 := right(str)

		if !assert.Equal(t, i2, i1) {
			break
		}
	}
}

func right(s string) int {
	if len(s) == 0 {
		return 0
	}
	str := manacherString(s)
	max := 0
	for i := 0; i < len(str); i++ {
		L := i - 1
		R := i + 1
		for L >= 0 && R < len(str) && str[L] == str[R] {
			L--
			R++
		}
		max = int(math.Max(float64(max), float64(R-L-1)))
	}
	return max / 2
}

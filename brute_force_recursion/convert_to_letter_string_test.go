package brute_force_recursion

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestNumber(t *testing.T) {
	N := 30
	testTimes := 50

	for i := 0; i < testTimes; i++ {
		length := rand.Intn(N + 1)
		s := randomString(length)
		fmt.Println(s)
		fmt.Println(Number(s))
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomString(length int) string {
	str := make([]string, length)
	for i := 0; i < len(str); i++ {
		str[i] = strconv.Itoa(rand.Intn(10))
	}
	return strings.Join(str, "")
}

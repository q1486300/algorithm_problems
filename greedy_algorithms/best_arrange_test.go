package greedy_algorithms

import (
	"github.com/stretchr/testify/assert"
	"math"
	"math/rand"
	"testing"
)

func TestBestArrange(t *testing.T) {
	testTimes := 1000000
	programSize := 12
	timeMax := 20

	for i := 0; i < testTimes; i++ {
		programs := generatePrograms(programSize, timeMax)

		b1 := BestArrange1(programs)
		b2 := BestArrange2(programs)

		if !assert.Equal(t, b1, b2) {
			break
		}
	}
}

func generatePrograms(programSize, timeMax int) Programs {
	ans := make(Programs, rand.Intn(programSize+1))
	for i := 0; i < len(ans); i++ {
		r1 := rand.Intn(timeMax + 1)
		r2 := rand.Intn(timeMax + 1)
		if r1 == r2 {
			ans[i] = NewProgram(r1, r1+1)
		} else {
			ans[i] = NewProgram(int(math.Min(float64(r1), float64(r2))),
				int(math.Max(float64(r1), float64(r2))))
		}
	}
	return ans
}

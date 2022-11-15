package binary_tree

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestMaxHappy(t *testing.T) {
	testTimes := 100000
	maxLevel := 4
	maxNexts := 7
	maxHappy := 100

	for i := 0; i < testTimes; i++ {
		boss := generateBoss(maxLevel, maxNexts, maxHappy)

		m1 := MaxHappy1(boss)
		m2 := MaxHappy2(boss)

		if !assert.Equal(t, m1, m2) {
			break
		}
	}
}

func generateBoss(maxLevel, maxNexts, maxHappy int) *Employee {
	if rand.Float32() < 0.02 {
		return nil
	}
	boss := NewEmployee(rand.Intn(maxHappy + 1))
	generateNexts(boss, 1, maxLevel, maxNexts, maxHappy)
	return boss
}

func generateNexts(e *Employee, level, maxLevel, maxNexts, maxHappy int) {
	if level > maxLevel {
		return
	}
	nextsSize := rand.Intn(maxNexts + 1)
	for i := 0; i < nextsSize; i++ {
		next := NewEmployee(rand.Intn(maxHappy + 1))
		e.nexts = append(e.nexts, next)
		generateNexts(next, level+1, maxLevel, maxNexts, maxHappy)
	}
}

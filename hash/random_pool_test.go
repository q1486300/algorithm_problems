package hash

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRandomPool(t *testing.T) {
	stringPool := NewPool[string]()
	stringPool.Insert("a")
	stringPool.Insert("b")
	stringPool.Insert("c")

	for i := 0; i < 5; i++ {
		fmt.Println(stringPool.GetRandom())
	}

	stringPool.Delete("b")

	fmt.Println("=============")

	for i := 0; i < 5; i++ {
		fmt.Println(stringPool.GetRandom())
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

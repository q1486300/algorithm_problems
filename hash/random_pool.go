package hash

import (
	"math/rand"
)

type Pool[K comparable] struct {
	keyIndexMap map[K]int
	indexKeyMap map[int]K
	size        int
}

func NewPool[K comparable]() *Pool[K] {
	return &Pool[K]{
		keyIndexMap: make(map[K]int),
		indexKeyMap: make(map[int]K),
		size:        0,
	}
}

func (p *Pool[K]) Insert(key K) {
	_, ok := p.keyIndexMap[key]
	if !ok {
		p.keyIndexMap[key] = p.size
		p.indexKeyMap[p.size] = key
		p.size++
	}
}

func (p *Pool[K]) Delete(key K) {
	deleteIndex, ok := p.keyIndexMap[key]
	if ok {
		p.size--
		lastIndex := p.size
		lastKey := p.indexKeyMap[lastIndex]
		p.keyIndexMap[lastKey] = deleteIndex
		p.indexKeyMap[deleteIndex] = lastKey
		delete(p.keyIndexMap, key)
		delete(p.indexKeyMap, lastIndex)
	}
}

func (p *Pool[K]) GetRandom() (result K) {
	if p.size == 0 {
		return
	}
	randomIndex := rand.Intn(p.size)
	result = p.indexKeyMap[randomIndex]
	return
}

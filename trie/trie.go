package trie

// 測試連結: https://leetcode.com/problems/implement-trie-ii-prefix-tree/
// 使用雜湊表實現前綴樹
type Node struct {
	pass  int
	end   int
	nexts map[int]*Node
}

func NewNode() *Node {
	return &Node{
		pass:  0,
		end:   0,
		nexts: make(map[int]*Node),
	}
}

type Trie struct {
	root *Node
}

func NewTrie() Trie {
	return Trie{root: NewNode()}
}

func (t Trie) Insert(word string) {
	chs := []rune(word)
	node := t.root
	node.pass++
	for _, ch := range chs {
		index := int(ch)
		_, ok := node.nexts[index]
		if !ok {
			node.nexts[index] = NewNode()
		}
		node = node.nexts[index]
		node.pass++
	}
	node.end++
}

func (t Trie) Erase(word string) {
	if t.CountWordsEqualTo(word) != 0 {
		chs := []rune(word)
		node := t.root
		node.pass--
		for _, ch := range chs {
			index := int(ch)
			node.nexts[index].pass--
			if node.nexts[index].pass == 0 {
				delete(node.nexts, index)
				return
			}
			node = node.nexts[index]
		}
		node.end--
	}
}

func (t Trie) CountWordsEqualTo(word string) int {
	chs := []rune(word)
	node := t.root
	for _, ch := range chs {
		index := int(ch)
		_, ok := node.nexts[index]
		if !ok {
			return 0
		}
		node = node.nexts[index]
	}
	return node.end
}

func (t Trie) CountWordsStartingWith(pre string) int {
	chs := []rune(pre)
	node := t.root
	for _, ch := range chs {
		index := int(ch)
		_, ok := node.nexts[index]
		if !ok {
			return 0
		}
		node = node.nexts[index]
	}
	return node.pass
}

package binary_tree

import "math"

type Employee struct {
	happy int
	nexts []*Employee
}

func NewEmployee(h int) *Employee {
	return &Employee{
		happy: h,
		nexts: make([]*Employee, 0),
	}
}

func MaxHappy1(boss *Employee) int {
	if boss == nil {
		return 0
	}
	return processMaxHappy1(boss, false)
}

// 目前來到的節點是 cur
// yes 表示 cur 的上級是否來
// 函數的涵義:
// 如果 yes 為 true，表示在 cur 上級已經確定來的情況下，cur 整棵樹能夠提供的最大快樂值是多少
// 如果 yes 為 false，表示在 cur 上級已經確定不來的情況下，cur 整棵樹能夠提供的最大快樂值是多少
func processMaxHappy1(cur *Employee, yes bool) int {
	if yes { // 如果 cur 的上級來的話，cur 沒得選，只能不來
		ans := 0
		for _, next := range cur.nexts {
			ans += processMaxHappy1(next, false)
		}
		return ans
	} else { // 如果 cur 的上級不來的話，cur 可以選，可以來也可以不來
		p1 := cur.happy
		p2 := 0
		for _, next := range cur.nexts {
			p1 += processMaxHappy1(next, true)
			p2 += processMaxHappy1(next, false)
		}
		return int(math.Max(float64(p1), float64(p2)))
	}
}

func MaxHappy2(head *Employee) int {
	allInfo := processMaxHappy2(head)
	return int(math.Max(float64(allInfo.yes), float64(allInfo.no)))
}

type MaxHappyInfo struct {
	yes, no int
}

func NewMaxHappyInfo(y, n int) MaxHappyInfo {
	return MaxHappyInfo{
		yes: y,
		no:  n,
	}
}

func processMaxHappy2(x *Employee) MaxHappyInfo {
	if x == nil {
		return NewMaxHappyInfo(0, 0)
	}
	yes := x.happy
	no := 0
	for _, next := range x.nexts {
		nextInfo := processMaxHappy2(next)
		yes += nextInfo.no
		no += int(math.Max(float64(nextInfo.yes), float64(nextInfo.no)))
	}
	return NewMaxHappyInfo(yes, no)
}

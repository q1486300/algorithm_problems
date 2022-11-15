package greedy_algorithms

import (
	"math"
	"sort"
)

type Program struct {
	start, end int
}

func NewProgram(start, end int) Program {
	return Program{
		start: start,
		end:   end,
	}
}

type Programs []Program

func (p Programs) Len() int {
	return len(p)
}

func (p Programs) Less(i, j int) bool {
	return p[i].end < p[j].end
}

func (p Programs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// 暴力! 所有情況都嘗試
func BestArrange1(programs Programs) int {
	if programs == nil || len(programs) == 0 {
		return 0
	}
	return processBestArrange(programs, 0, 0)
}

// 還剩下的會議都放在 programs 裡
// done 之前已經安排了多少會議的數量
// timeLine 目前來到的時間點是什麼

// 目前來到 timeLine 的時間點，已經安排了 done 個會議，剩下的會議 programs 可以自由安排
// 返回能安排的最多會議數量
func processBestArrange(programs Programs, done, timeLine int) int {
	if len(programs) == 0 {
		return done
	}
	// 還剩下會議
	max := done
	// 目前安排的會議是什麼會，每一個都列舉
	for i, program := range programs {
		if program.start >= timeLine {
			next := copyButExcept(programs, i)
			max = int(math.Max(float64(max), float64(processBestArrange(next, done+1, program.end))))
		}
	}
	return max
}

func copyButExcept(programs Programs, i int) Programs {
	ans := make(Programs, len(programs)-1)
	index := 0
	for k, program := range programs {
		if k != i {
			ans[index] = program
			index++
		}
	}
	return ans
}

// 會議的開始時間和結束時間都是數值，不會 < 0
func BestArrange2(programs Programs) int {
	sort.Sort(programs)
	timeLine := 0
	result := 0
	// 依序遍歷每個會議，結束時間早的會議先遍歷
	for _, program := range programs {
		if timeLine <= program.start {
			result++
			timeLine = program.end
		}
	}
	return result
}

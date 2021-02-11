package main

import (
	"fmt"
	"sort"
	"sync"
)

func main() {
	fmt.Println(minAbsDifference([]int{5, -7, 3, 5}, 6))
	fmt.Println(minAbsDifference([]int{7, -9, 15, -2}, -5))
	fmt.Println(minAbsDifference([]int{1, 2, 3}, -7))
	fmt.Println(minAbsDifference([]int{3530, -1549, 6835, -587, 3787, -1033, 4205, 1006, 5918, -2940, 6101, 3169, 3930, -7006, -7889, -5758, -3246, -5098, -2489, -9144, -6617, -1703, -4898, 5721, -6758, 3078, -3859, -9902, -7079, 4014, -8334, 8009}, 842213514))
	fmt.Println(minAbsDifference([]int{-7074297, 3076735, -5846354, 5008659, -126683, 7039557, 6708811, 3189666, -6102417, 6078975, -6448946, -4995910, 2964239, -3248847, -4392269, 7473223, -1356059, 3978911, 8009187, -316441, 6524770, 8280309, -2798383, 1310839, 6306594, -6548611, -9712711, 1639314}, 493409180))
}

func minabs(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	if a < b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func minAbsDifference(nums []int, goal int) int {
	var wg sync.WaitGroup
	wg.Add(2)
	var a, b []int
	go func() { a = list(nums[:len(nums)/2]); wg.Done() }()
	go func() { b = list(nums[len(nums)/2:]); wg.Done() }()
	wg.Wait()
	mn := abs(goal)
	for _, i := range a {
		g := goal - i
		j := sort.Search(len(b), func(i int) bool {
			return b[i] >= g
		})
		if j > 0 {
			mn = minabs(mn, g-b[j-1])
		}
		if j < len(b) {
			mn = minabs(mn, g-b[j])
		}
	}
	return mn
}

func list(n []int) []int {
	var r = []int{0}
	for _, v := range n {
		for _, i := range r {
			r = append(r, i+v)
		}
	}
	sort.Ints(r)
	return r
}

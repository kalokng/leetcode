package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func twoSum(nums []int, target int) []int {
	idx := make([]int, len(nums))
	for i := range idx {
		idx[i] = i
	}
	sort.Slice(idx, func(i, j int) bool {
		return nums[idx[i]] < nums[idx[j]]
	})
	for i, v := range idx {
		remain := target - nums[v]
		s := sort.Search(len(idx), func(n int) bool { return nums[idx[n]] >= remain })
		if s == i {
			s++
		}
		if s >= len(idx) {
			continue
		}
		if nums[idx[s]] != remain {
			continue
		}
		return []int{idx[i], idx[s]}
	}
	panic("not found")
}

func main() {
	s := make([]int, 64)
	for i := range s {
		s[i] = -(1 << i)
	}
	s[len(s)-1] = 0
	rand.Shuffle(len(s), func(i, j int) { s[i], s[j] = s[j], s[i] })
	fmt.Println(s)

	for i := 0; i < 100; i++ {
		x := rand.Intn(64)
		y := rand.Intn(64)
		t := s[x] + s[y]
		a := twoSum(s, t)
		fmt.Println(t, a, s[a[0]]+s[a[1]])
	}
	fmt.Println(twoSum([]int{3, 3, 2, 2}, 4))
}

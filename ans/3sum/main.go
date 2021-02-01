package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{-1, -1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{3, 0, -2, -1, 1, 2}))
}

func threeSum(nums []int) [][]int {
	m := make(map[int]int)
	var set = make(map[[3]int]struct{})
	for _, v := range nums {
		for k, n := range m {
			r := -v - k
			if r == k {
				if n > 1 {
					l := [3]int{v, r, k}
					sort.Ints(l[:])
					set[l] = struct{}{}
				}
				continue
			}
			if m[r] > 0 {
				l := [3]int{v, r, k}
				sort.Ints(l[:])
				set[l] = struct{}{}
			}
		}
		m[v]++
	}
	var res [][]int
	for k := range set {
		l := make([]int, 3)
		copy(l, k[:])
		res = append(res, l)
	}
	return res
}

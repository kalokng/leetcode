package main

import "fmt"

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	fmt.Println(searchRange([]int{}, 0))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 10}, 8))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 9, 10}, 8))
}

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	l, r := 0, len(nums)-1
	le, re := l, r
	for l < r {
		m := (l + r) / 2
		v := nums[m]
		switch {
		case v < target:
			l = m + 1
		case v > target:
			r = m - 1
			re = r
		default:
			r = m
			le = m
		}
	}
	if nums[l] != target {
		return []int{-1, -1}
	}
	if l > le {
		le = l
	}
	for le < re {
		m := (le + re + 1) / 2
		v := nums[m]
		switch {
		case v < target:
			le = m + 1
		case v > target:
			re = m - 1
		default:
			le = m
		}
	}
	return []int{l, re}
}

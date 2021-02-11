package main

import (
	"fmt"
)

func main() {
	fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 0))
	fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 3))
	fmt.Println(search([]int{0, 1, 2, 3}, 2))
	fmt.Println(search([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1}, 2))
	fmt.Println(search([]int{2, 2, 2, 0, 0, 1}, 0))
}

func search(nums []int, target int) bool {
	return s(nums, target, 0, len(nums)-1) >= 0
}

func s(nums []int, target int, l, r int) int {
	fmt.Println(l, r)
	if r-l <= 1 {
		for i := l; i <= r; i++ {
			if nums[i] == target {
				return i
			}
		}
		return -1
	}

	if nums[l] == nums[r] {
		if nums[l] == target {
			return l
		}
		m := (l + r) / 2
		if nums[m] == nums[l] {
			if n := s(nums, target, l+1, m-1); n > 0 {
				return n
			}
			return s(nums, target, m+1, r-1)
		}
		if nums[m] == target {
			return m
		}
		if nums[m] > nums[l] {
			if nums[m] > target && nums[l] < target {
				return s(nums, target, l+1, m-1)
			} else {
				return s(nums, target, m+1, r-1)
			}
		} else {
			if nums[m] < target && target < nums[r] {
				return s(nums, target, m+1, r-1)
			} else {
				return s(nums, target, l+1, m-1)
			}
		}
	}

	if nums[l] < nums[r] {
		i, j := l, r
		// normal array
		for {
			m := (i + j) / 2
			switch {
			case nums[i] > target, nums[j] < target:
				return -1
			case nums[m] > target:
				j = m - 1
			case nums[m] < target:
				i = m + 1
			default:
				return m
			}
		}
	}

	switch {
	case nums[l] > target:
		// target is somewhere in the 2nd half
		m := (l + r) / 2
		v := nums[m]
		switch {
		case v >= nums[l], v < target:
			// target in between (m,r]
			return s(nums, target, m+1, r)
		case v > target:
			// target in 2nd half of (0,m)
			return s(nums, target, 1, m-1)
		default:
			return m
		}
	case nums[l] < target:
		// target is somewhere in the 1st half
		m := (l + r) / 2
		v := nums[m]
		switch {
		case v <= nums[l], v > target:
			// target in between (0,m)
			return s(nums, target, 1, m-1)
		case v < target:
			// target in 1st half of (m,r]
			return s(nums, target, m+1, r)
		default:
			return m
		}
	}
	return l
}

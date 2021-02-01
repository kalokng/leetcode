package main

import "fmt"

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 6))
	fmt.Println(search([]int{0, 1, 2, 4, 5, 6, 7}, 6))
	fmt.Println(search([]int{0, 1, 2, 3}, 2))
}

func search(nums []int, target int) int {
	return s(nums, target, 0, len(nums)-1)
}

func s(nums []int, target int, l, r int) int {
	if r-l <= 1 {
		for i := l; i <= r; i++ {
			if nums[i] == target {
				return i
			}
		}
		return -1
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
		case v > nums[l], v < target:
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
		case v < nums[l], v > target:
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

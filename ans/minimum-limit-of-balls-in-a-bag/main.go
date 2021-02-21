package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minimumSize([]int{9}, 1))
	fmt.Println(minimumSize([]int{9}, 2))
	fmt.Println(minimumSize([]int{7, 17}, 2))
	fmt.Println(minimumSize([]int{9}, 3))
	fmt.Println(minimumSize([]int{9}, 4))
	fmt.Println(minimumSize([]int{2, 4, 8, 2}, 4))
	fmt.Println(minimumSize([]int{2, 4, 8, 1}, 4))
	fmt.Println(minimumSize([]int{2, 4, 10, 1, 1}, 4))
	fmt.Println(minimumSize([]int{2, 4, 10, 1}, 4))
	fmt.Println(minimumSize([]int{431, 922, 158, 60, 192, 14, 788, 146, 788, 775, 772, 792, 68, 143, 376, 375, 877, 516, 595, 82, 56, 704, 160, 403, 713, 504, 67, 332, 26}, 80))
	//129
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minimumSize(nums []int, maxOperations int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	i := sort.Search(len(nums), func(i int) bool {
		size := nums[i]
		n := 0
		for j := 0; j < i; j++ {
			n += nums[j]/size - 1
			if nums[j]%size > 0 {
				n++
			}
			if n >= maxOperations {
				return true
			}
		}
		return n >= maxOperations
	})
	var mn, mx int
	if i < len(nums) {
		mn, mx = nums[i], nums[i-1]
	} else {
		mn, mx = 1, nums[i-1]
	}

	fmt.Println(mn, mx)
	return sort.Search(mx-mn+1, func(v int) bool {
		v += mn
		n := 0
		fmt.Println(v)
		for j := 0; j < i; j++ {
			if nums[j] < v {
				continue
			}
			n += nums[j]/v - 1
			if nums[j]%v > 0 {
				n++
			}
			if n > maxOperations {
				return false
			}
		}
		return n <= maxOperations
	}) + mn
}

package main

import "fmt"

func main() {
	for _, nums := range [][]int{
		[]int{1, 1, 1, 2, 2, 3},
		[]int{0, 0, 1, 1, 1, 1, 2, 3, 3},
	} {
		n := removeDuplicates(nums)
		fmt.Println(nums[:n])
	}
}

func removeDuplicates(nums []int) int {
	i := 0
	n := 1
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			n = 1
			i++
			nums[i] = nums[j]
			continue
		}
		if n < 2 {
			i++
			nums[i] = nums[j]
		}
		n++
		continue
	}
	return i + 1
}

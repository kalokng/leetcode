package main

import "fmt"

func main() {
	fmt.Println(canChoose([][]int{{1, -1, -1}, {3, -2, 0}}, []int{1, -1, 0, 1, -1, -1, 3, -2, 0}))
	fmt.Println(canChoose([][]int{{1, -1, -1}, {3, -2, 0}}, []int{1, -1, 0, 1, -1, -1, 3, -2, 0}))
}

func canChoose(groups [][]int, nums []int) bool {
	j := 0
	for i := 0; i < len(nums); i++ {
		k := 0
		if len(nums)-i < len(groups[j]) {
			return false
		}
		for ; k < len(groups[j]); k++ {
			if groups[j][k] != nums[i+k] {
				break
			}
		}
		if k == len(groups[j]) {
			j++
			if j >= len(groups) {
				return true
			}
			i += k - 1
		}
	}
	return false
}

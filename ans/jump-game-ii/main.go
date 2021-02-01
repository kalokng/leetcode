package main

import "fmt"

func jump(nums []int) int {
	step := 0
	reach := 0
	next := 0
	for i, v := range nums {
		if i > reach {
			step++
			reach = next
		}
		if v+i > next {
			next = v + i
		}
	}
	return step
}

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
	fmt.Println(jump([]int{2, 3, 0, 1, 4}))
	fmt.Println(jump([]int{3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3}))
}

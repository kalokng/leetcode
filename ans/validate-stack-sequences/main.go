package main

import "fmt"

func main() {
	fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}))
	fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 2, 1}))
	fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}))
}

func validateStackSequences(pushed []int, popped []int) bool {
	s, i := 0, 0
	for _, v := range popped {
		if s > 0 && pushed[s-1] == v {
			s--
			continue
		}
		if i >= len(pushed) {
			return false
		}
		for ; i < len(pushed); i++ {
			if pushed[i] == v {
				i++
				break
			}
			pushed[s] = pushed[i]
			s++
		}
	}
	return true
}

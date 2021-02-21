package main

import "fmt"

func main() {
	fmt.Println(minOperations("001011"))
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

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func minOperations(boxes string) []int {
	var b []int
	for i := 0; i < len(boxes); i++ {
		if boxes[i] == '1' {
			b = append(b, i)
		}
	}
	var r = make([]int, len(boxes))
	for i := range r {
		s := 0
		for _, v := range b {
			s += abs(v - i)
		}
		r[i] = s
	}
	return r
}

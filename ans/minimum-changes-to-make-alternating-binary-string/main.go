package main

import "fmt"

func main() {
	fmt.Println(minOperations("0100"))
	fmt.Println(minOperations("10"))
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

func minOperations(s string) int {
	var c = make([]int, 2)
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '0':
			c[i%2]++
		case '1':
			c[(i+1)%2]++
		}
	}
	return min(c[0], c[1])
}

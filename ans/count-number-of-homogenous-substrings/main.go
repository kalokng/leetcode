package main

import "fmt"

func main() {
	fmt.Println(countHomogenous("abbcccaa"))
	fmt.Println(countHomogenous("xy"))
	fmt.Println(countHomogenous("zzzzz"))
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

const mod = 1e9 + 7

func countHomogenous(s string) int {
	var l int
	n := 1
	for i := 1; i < len(s); i++ {
		if s[i] != s[l] {
			l = i
		}
		n = (n + i - l + 1) % mod
	}
	return n
}

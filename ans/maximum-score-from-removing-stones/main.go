package main

import "fmt"

func main() {
	fmt.Println(maximumScore(2, 4, 6))
	fmt.Println(maximumScore(4, 4, 6))
	fmt.Println(maximumScore(1, 8, 8))
	fmt.Println(maximumScore(1, 8, 11))
	fmt.Println(maximumScore(1, 3, 11))
}

func maximumScore(a int, b int, c int) int {
	if a > b {
		a, b = b, a
	}
	if a > c {
		a, b, c = c, a, b
	} else if b > c {
		b, c = c, b
	}
	d := c - b
	if d >= a {
		return a + b
	}
	sum := (b + c - a)
	b = sum / 2
	c = sum - b
	return a + b
}

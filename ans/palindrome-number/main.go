package main

import "fmt"

func main() {
	for _, v := range []int{
		121, -121, 10, 100, 101, 0, 1, 11, 111, 1111,
		2147483648,
		2147447412,
	} {
		fmt.Println(v, isPalindrome(v))
	}
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x%10 == 0 {
		return x == 0
	}
	y := 0
	for y < x {
		x, y = x/10, 10*y+x%10
	}
	if y > x {
		y /= 10
	}
	return y == x
}

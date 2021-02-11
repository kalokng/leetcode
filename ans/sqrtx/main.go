package main

import "fmt"

func main() {
	fmt.Println(mySqrt(0))
	fmt.Println(mySqrt(4))
	fmt.Println(mySqrt(8))
	fmt.Println(mySqrt(9))
	fmt.Println(mySqrt(2093144237))
}

func mySqrt(x int) int {
	r := x
	for r*r > x {
		r = (r + x/r) / 2
	}
	return r
}

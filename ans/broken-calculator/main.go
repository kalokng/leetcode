package main

import "fmt"

func main() {
	fmt.Println(brokenCalc(2, 3))
	fmt.Println(brokenCalc(5, 8))
	fmt.Println(brokenCalc(3, 10))
	fmt.Println(brokenCalc(1024, 1))
	fmt.Println(brokenCalc(9, 10))
	fmt.Println(brokenCalc(8, 9))
}

func brokenCalc(X int, Y int) int {
	if Y <= X {
		return X - Y
	}
	var n int
	for X < Y {
		X *= 2
		n++
	}
	d := X - Y
	s := n
	for d > 0 && n > 0 {
		s += d % 2
		n--
		d /= 2
	}
	return d + s
}

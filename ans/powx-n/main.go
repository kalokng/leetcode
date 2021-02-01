package main

import "fmt"

func main() {
	fmt.Println(myPow(2.0, 10))
	fmt.Println(myPow(2.1, 3))
	fmt.Println(myPow(2.0, -2))
	fmt.Println(myPow(1.01, 20))
	fmt.Println(myPow(1.00001, 200000))
	fmt.Println(myPow(1.000002, 2000000))
	fmt.Println(myPow(1.000002, -2000000))
}

func myPow(x float64, n int) float64 {
	if n < 0 {
		n = -n
		x = 1 / x
	}
	r := 1.0
	for n > 0 {
		if n&1 != 0 {
			r *= x
		}
		x *= x
		n >>= 1
	}
	return r
}

package main

import "fmt"

const max = 1 << 31

func divide(dividend int, divisor int) int {
	if divisor == 1 {
		return dividend
	}
	if divisor == -1 {
		if dividend == -max {
			return max - 1
		}
		return -dividend
	}
	if dividend == 0 {
		return 0
	}
	var sign = 1
	if dividend < 0 {
		sign = -sign
		dividend = -dividend
	}
	if divisor < 0 {
		sign = -sign
		divisor = -divisor
	}
	var bit []int
	for divisor <= dividend {
		bit = append(bit, divisor)
		divisor = divisor + divisor
	}
	var val int
	for l := len(bit) - 1; l >= 0 && dividend > 0; l-- {
		if dividend >= bit[l] {
			val = val | (1 << l)
			dividend -= bit[l]
		}
	}
	return sign * val
}

func main() {
	fmt.Println(divide(10, 3))
	fmt.Println(divide(7, -3))
	fmt.Println(divide(0, 1))
	fmt.Println(divide(1, 1))
	fmt.Println(divide(2147483647, 1))
	fmt.Println(divide(-2147483648, -1))
	fmt.Println(divide(333, 7))
	fmt.Println(divide(-10, 3))
	fmt.Println(divide(-7, -3))
	fmt.Println(divide(0, -1))
	fmt.Println(divide(-1, 1))
	fmt.Println(divide(1, -1))
	fmt.Println(divide(-1, -1))
	fmt.Println(divide(-2147483648, 1))
	fmt.Println(divide(2147483647, -1))
	fmt.Println(divide(-333, 7))
}

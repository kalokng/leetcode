package main

import "fmt"

const max = 1 << 31

func reverse(x int) int {
	if x < 0 {
		return -r(uint32(-x), max)
	}
	return r(uint32(x), max-1)
}

func r(x uint32, max uint32) int {
	var y uint32
	for x > 0 {
		var c uint32
		x, c = x/10, x%10
		if y != 0 && (max-c)/y < 10 {
			return 0
		}
		y = 10*y + c
	}
	return int(y)
}

func main() {
	fmt.Println(max)
	fmt.Println(reverse(76723))
	fmt.Println(reverse(86723))
	fmt.Println(reverse(-86723))
	fmt.Println(reverse(2147483648))
	fmt.Println(reverse(2147447412))
	fmt.Println(reverse(max - 1))
	fmt.Println(reverse(-max))
}

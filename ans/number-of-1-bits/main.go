package main

import "fmt"

func main() {
	fmt.Println(hammingWeight(0b00000000000000000000000000001011))
	fmt.Println(hammingWeight(0b00000000000000000000000010000000))
	fmt.Println(hammingWeight(0b11111111111111111111111111111101))
}

func hammingWeight(num uint32) int {
	n := 0
	for num > 0 {
		n += int(num & 1)
		num >>= 1
	}
	return n
}

package main

import "fmt"

func main() {
	fmt.Println(concatenatedBinary(1))
	fmt.Println(concatenatedBinary(3))
	fmt.Println(concatenatedBinary(12))
	fmt.Println(concatenatedBinary(5e4))
}

const mod = 1e9 + 7

func concatenatedBinary(n int) int {
	var v int
	i := 1
	b := 1
	for s := 2; s <= n; s <<= 1 {
		for ; i < s; i++ {
			v = (v<<b + i) % mod
		}
		b++
	}
	for ; i <= n; i++ {
		v = (v<<b + i) % mod
	}
	return v
}

package main

import "fmt"

func main() {
	fmt.Println(getSum(1, 2))
	fmt.Println(getSum(-2, 3))
	fmt.Println(getSum(-2, -3))
}

const uintSize = 32 << (^uint(0) >> 32 & 1) // 32 or 64
const maxUint = 1 << (uintSize - 1)

type pair struct {
	b bool
	c uint
}

var addbit = [8]pair{
	{false, 0},
	{true, 0},
	{true, 0},
	{false, 4},
	{true, 0},
	{false, 4},
	{false, 4},
	{true, 4},
}

func getSum(a int, b int) int {
	var c uint
	var v uint
	ua, ub := uint(a), uint(b)
	for i := uint(1); i != 0; i = i << 1 {
		var ba, bb uint
		if ua&i > 0 {
			ba = 1
		}
		if ub&i > 0 {
			bb = 2
		}
		p := addbit[ba|bb|c]
		c = p.c
		if p.b {
			v = v | i
		}
	}
	return int(v)
}

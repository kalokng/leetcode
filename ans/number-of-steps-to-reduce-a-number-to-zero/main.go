package main

import "fmt"

func main() {
	fmt.Println(numberOfSteps(14))
	fmt.Println(numberOfSteps(8))
	fmt.Println(numberOfSteps(123))
}

func numberOfSteps(num int) int {
	n := 0
	for num > 1 {
		n += num%2 + 1
		num /= 2
	}
	return n + num
}

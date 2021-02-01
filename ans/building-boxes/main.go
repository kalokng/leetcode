package main

import "fmt"

func main() {
	fmt.Println(minimumBoxes(3))
	fmt.Println(minimumBoxes(4))
	fmt.Println(minimumBoxes(10))
	fmt.Println(minimumBoxes(11))
	fmt.Println(minimumBoxes(12))
	fmt.Println(minimumBoxes(13))
	fmt.Println(minimumBoxes(14))
	fmt.Println(minimumBoxes(15))
	fmt.Println(minimumBoxes(16))
	fmt.Println(minimumBoxes(17))
	fmt.Println(minimumBoxes(18))
	fmt.Println(minimumBoxes(19))
	fmt.Println(minimumBoxes(20))
	fmt.Println(minimumBoxes(21))
}

func minimumBoxes(n int) int {
	b := 0
	ml := 1
	t := 0
	for t < n {
		b += ml
		ml++
		t += b
	}
	n -= t - b
	b -= ml - 1
	i := 1
	for n > 0 {
		n -= i
		b++
		i++
	}
	return b
}

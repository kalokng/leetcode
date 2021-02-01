package main

import "fmt"

func main() {
	fmt.Println(getSmallestString(3, 27))
	fmt.Println(getSmallestString(5, 73))
	fmt.Println(getSmallestString(5, 130))
}

func getSmallestString(n int, k int) string {
	k -= n
	b := make([]byte, n)
	i := n - 1
	for ; k >= 25; i-- {
		b[i] = 25
		k -= 25
	}
	if k > 0 {
		b[i] = byte(k)
	}
	for i := range b {
		b[i] += 'a'
	}
	return string(b)
}

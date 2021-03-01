// +build ignore

package main

import "fmt"

func main() {
	fmt.Println(findLongestWord("abpcplea", []string{"ale", "apple", "monkey", "plea"}))
	fmt.Println(findLongestWord("abpcplea", []string{"a", "b", "c"}))
}

func findLongestWord(s string, d []string) string {
	n := make([]int, len(d))
	for i := 0; i < len(s); i++ {
		c := s[i]
		for j, k := range n {
			if k >= len(d[j]) {
				continue
			}
			if c == d[j][k] {
				n[j]++
			}
		}
	}
	mx := ""
	for j, k := range n {
		if k == len(d[j]) && len(mx) <= k {
			if len(mx) == k && mx < d[j] {
				continue
			}
			mx = d[j]
		}
	}
	return mx
}

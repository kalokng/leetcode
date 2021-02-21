package main

import "fmt"

func main() {
	fmt.Println(mergeAlternately("abc", "pqr"))
	fmt.Println(mergeAlternately("ab", "pqrs"))
	fmt.Println(mergeAlternately("abcd", "pq"))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func mergeAlternately(word1 string, word2 string) string {
	var s []byte
	l := len(word1)
	if l > len(word2) {
		l = len(word2)
	}
	for i := 0; i < l; i++ {
		s = append(s, word1[i], word2[i])
	}
	return string(s) + word1[l:] + word2[l:]
}

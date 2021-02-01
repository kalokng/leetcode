package main

import (
	"fmt"
)

func main() {
	fmt.Println(closeStrings("cabbba", "abbccc"))
	fmt.Println(closeStrings("abc", "bca"))
	fmt.Println(closeStrings("aa", "a"))
	fmt.Println(closeStrings("cabbba", "aabbss"))
}

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	w1 := make(map[byte]int)
	w2 := make(map[byte]int)
	for i := 0; i < len(word1); i++ {
		w1[word1[i]]++
	}
	for i := 0; i < len(word2); i++ {
		w2[word2[i]]++
	}

	m := make(map[byte]struct{})
	c := make(map[int]int)
	for k, v := range w1 {
		m[k] = struct{}{}
		c[v]++
	}
	for k, v := range w2 {
		if _, ok := m[k]; !ok {
			return false
		}
		delete(m, k)
		c[v]--
	}
	if len(m) > 0 {
		return false
	}
	for _, v := range c {
		if v != 0 {
			return false
		}
	}
	return true
}

package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("cacb", "cbba"))
	fmt.Println(longestPalindrome("ab", "ab"))
	fmt.Println(longestPalindrome("aa", "bb"))
	fmt.Println(longestPalindrome("fbb", "eba"))
	fmt.Println(longestPalindrome("ceebeddc", "d"))
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

func longestPalindrome(word1 string, word2 string) int {
	w1 := map[byte]struct{}{}
	for i := 0; i < len(word1); i++ {
		w1[word1[i]] = struct{}{}
	}
	done := false
	for i := 0; i < len(word2); i++ {
		if _, ok := w1[word2[i]]; ok {
			done = true
			break
		}
	}
	if !done {
		return 0
	}
	s := word1 + word2
	c := make(map[P]int)
	c2 := make(map[int]P)
	lp(c, c2, s, 0, len(word1)+len(word2)-1, len(word1))
	mx := 0
	for k := range c2 {
		if k > mx {
			mx = k
		}
	}
	return mx
}

type P struct {
	i, j int
}

func lp(c map[P]int, c2 map[int]P, s string, i, j int, l1 int) (res int) {
	if i >= j {
		return j - i + 1
	}
	p := P{i, j}
	if v, ok := c[p]; ok {
		return v
	}
	defer func() {
		c[p] = res
	}()

	if s[i] == s[j] {
		res = 2 + lp(c, c2, s, i+1, j-1, l1)
		if i < l1 && j >= l1 {
			c2[res] = p
		}
		return res
	}
	return max(lp(c, c2, s, i, j-1, l1), lp(c, c2, s, i+1, j, l1))
}

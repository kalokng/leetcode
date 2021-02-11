package main

import "fmt"

func main() {
	fmt.Println(largestMerge("cabaa", "bcaaa"))
	fmt.Println(largestMerge("abcabc", "abdcaba"))
}

func largestMerge(word1 string, word2 string) string {
	i := 0
	j := 0
	c1 := word1[i]
	c2 := word2[j]
	var res []byte
	for {
		if word1[i:] > word2[j:] {
			res = append(res, c1)
			i++
			if i >= len(word1) {
				break
			}
			c1 = word1[i]
		} else {
			res = append(res, c2)
			j++
			if j >= len(word2) {
				break
			}
			c2 = word2[j]
		}
	}
	return string(res) + word1[i:] + word2[j:]
}

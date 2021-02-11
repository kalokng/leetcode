package main

import "fmt"

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
	fmt.Println(lengthOfLastWord(" "))
	fmt.Println(lengthOfLastWord(""))
	fmt.Println(lengthOfLastWord("a"))
	fmt.Println(lengthOfLastWord("abcde"))
	fmt.Println(lengthOfLastWord("Hello World abc"))
}

func lengthOfLastWord(s string) int {
	j := len(s) - 1
	for ; j >= 0; j-- {
		if s[j] != ' ' {
			break
		}
	}
	i := j
	for ; i >= 0; i-- {
		if s[i] == ' ' {
			break
		}
	}
	return j - i
}

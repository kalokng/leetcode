package main

import "fmt"

func main() {
	fmt.Println(exist([][]byte{
		[]byte("ABCE"),
		[]byte("SFCS"),
		[]byte("ADEE"),
	}, "ABCCED"))
	fmt.Println(exist([][]byte{
		[]byte("ABCE"),
		[]byte("SFCS"),
		[]byte("ADEE"),
	}, "SEE"))
	fmt.Println(exist([][]byte{
		[]byte("ABCE"),
		[]byte("SFCS"),
		[]byte("ADEE"),
	}, "ABCB"))
	fmt.Println(exist([][]byte{
		[]byte("CAA"),
		[]byte("AAA"),
		[]byte("BCD"),
	}, "AAB"))
}

func exist(board [][]byte, word string) bool {
	m, n := len(board)-1, len(board[0])-1
	c := word[0]

	for i := range board {
		for j := range board[i] {
			if c != board[i][j] {
				continue
			}
			if search(board, word, m, n, i, j, 1) {
				return true
			}
		}
	}
	return false
}

type pair struct {
	i, j int
}

func search(b [][]byte, word string, m, n int, i, j int, k int) bool {
	if len(word) == k {
		return true
	}
	b[i][j] = 0
	defer func() { b[i][j] = word[k-1] }()
	f := func(i, j int) bool {
		if b[i][j] != word[k] {
			return false
		}
		return search(b, word, m, n, i, j, k+1)
	}
	return (i > 0 && f(i-1, j)) ||
		(i < m && f(i+1, j)) ||
		(j > 0 && f(i, j-1)) ||
		(j < n && f(i, j+1))
}

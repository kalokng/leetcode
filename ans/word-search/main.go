// +build ignore

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
}

func exist(board [][]byte, word string) bool {
	m, n := len(board)-1, len(board[0])-1
	v := make([][]bool, m+3)
	for i := 0; i < m+3; i++ {
		v[i] = make([]bool, n+3)
		v[i][0] = true
		v[i][n+2] = true
	}
	for j := 1; j < n+2; j++ {
		v[0][j] = true
		v[m+2][j] = true
	}
	c := word[0]

	for i := range board {
		for j := range board[i] {
			if c != board[i][j] {
				continue
			}
			v[i+1][j+1] = true
			if search(board, word, m, n, i, j, v, 1) {
				return true
			}
			v[i+1][j+1] = false
		}
	}
	return false
}

type pair struct {
	i, j int
}

func search(b [][]byte, word string, m, n int, i, j int, v [][]bool, k int) bool {
	if len(word) == k {
		return true
	}
	var q []pair
	if !v[i][j+1] && b[i-1][j] == word[k] {
		q = append(q, pair{i - 1, j})
	}
	if !v[i+2][j+1] && b[i+1][j] == word[k] {
		q = append(q, pair{i + 1, j})
	}
	if !v[i+1][j] && b[i][j-1] == word[k] {
		q = append(q, pair{i, j - 1})
	}
	if !v[i+1][j+2] && b[i][j+1] == word[k] {
		q = append(q, pair{i, j + 1})
	}
	for _, p := range q {
		v[p.i+1][p.j+1] = true
		if search(b, word, m, n, p.i, p.j, v, k+1) {
			return true
		}
		v[p.i+1][p.j+1] = false
	}
	return false
}

package main

import (
	"fmt"
	"math/bits"
	"sort"
)

type Choice map[Pos]uint16

func (c Choice) Pos(i, j int) uint16 {
	return c[pos(i, j)]
}

func solveSudoku(board [][]byte) {
	if done(board) {
		return
	}
	choice := initBoard(board)

update:
	oldChoice := make(Choice)
	for pp, opt := range choice {
		if val, ok := hit[opt]; ok {
			delete(choice, pp)
			board[pp.X][pp.Y] = val
			updateChoice(choice, oldChoice, pp, val)
			goto update
		}
	}

	lvl := 1
	for {
		if solveBoard(board, choice, lvl) {
			fmt.Println(lvl)
			return
		}
		lvl++
	}
}

func solvable(choice Choice) bool {
	for _, v := range choice {
		if v == 0 {
			return false
		}
	}
	return true
}

func solveBoard(board [][]byte, choice Choice, lvl int) bool {
	c := make([]Pos, 0, len(choice))
	for p := range choice {
		c = append(c, p)
	}
	sort.Slice(c, func(i, j int) bool {
		return bits.OnesCount16(choice[c[i]]) < bits.OnesCount16(choice[c[j]])
	})

	for _, p := range c {
		opt := choice[p]
		for v := byte('1'); v <= '9'; v++ {
			if opt&val[v] == 0 {
				continue
			}
			if tryBoard(board, choice, p, v, lvl) {
				return true
			}
		}
	}
	return len(choice) == 0
}

func tryBoard(board [][]byte, choice Choice, p Pos, v byte, lvl int) (ok bool) {
	oldChoice := Choice{p: choice[p]}
	var oldBoard []Pos
	defer func() {
		if ok {
			return
		}
		for p, v := range oldChoice {
			choice[p] = v
		}
		for _, p := range oldBoard {
			board[p.X][p.Y] = '.'
		}
	}()

update:
	delete(choice, p)
	board[p.X][p.Y] = v
	oldBoard = append(oldBoard, p)
	updateChoice(choice, oldChoice, p, v)
	for pp, opt := range choice {
		if opt == 0 {
			return false
		}
		if val, ok := hit[opt]; ok {
			v = val
			p = pp
			goto update
		}
	}

	if lvl == 0 {
		return false
	}
	return solveBoard(board, choice, lvl-1)
}

func updateChoice(choice, oldChoice Choice, p Pos, v byte) {
	choose := func(pp Pos) {
		if opt, ok := choice[pp]; ok && (opt&val[v] != 0) {
			if _, ok := oldChoice[pp]; !ok {
				oldChoice[pp] = opt
			}
			choice[pp] = opt &^ val[v]
		}
	}
	for i := byte(0); i < 9; i++ {
		choose(Pos{p.X, i})
		choose(Pos{i, p.Y})
	}
	cellr := 3 * (p.X / 3)
	cellc := 3 * (p.Y / 3)
	for x := byte(0); x < 3; x++ {
		for y := byte(0); y < 3; y++ {
			choose(Pos{cellr + x, cellc + y})
		}
	}
}

const full uint16 = 1<<9 - 1

var val = map[byte]uint16{
	'1': 1,
	'2': 1 << 1,
	'3': 1 << 2,
	'4': 1 << 3,
	'5': 1 << 4,
	'6': 1 << 5,
	'7': 1 << 6,
	'8': 1 << 7,
	'9': 1 << 8,
}
var hit = map[uint16]byte{
	1:      '1',
	1 << 1: '2',
	1 << 2: '3',
	1 << 3: '4',
	1 << 4: '5',
	1 << 5: '6',
	1 << 6: '7',
	1 << 7: '8',
	1 << 8: '9',
}

type Pos struct {
	X, Y byte
}

func pos(i, j int) Pos {
	return Pos{byte(i), byte(j)}
}

func initBoard(board [][]byte) Choice {
	choice := make(Choice)
	for i, r := range board {
		for j, c := range r {
			if c != '.' {
				continue
			}
			v := full
			for x := 0; x < 9; x++ {
				v &= ^val[board[i][x]]
				v &= ^val[board[x][j]]
			}
			cellr := 3 * (i / 3)
			cellc := 3 * (j / 3)
			for x := 0; x < 3; x++ {
				for y := 0; y < 3; y++ {
					v &= ^val[board[x+cellr][y+cellc]]
				}
			}
			choice[pos(i, j)] = v
		}
	}
	return choice
}

func done(b [][]byte) bool {
	for _, r := range b {
		for _, c := range r {
			if c == '.' {
				return false
			}
		}
	}
	return true
}

func main() {
	//b := [][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
	//	{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
	//	{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
	//	{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
	//	{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
	//	{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
	//	{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
	//	{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
	//	{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	//b := [][]byte{
	//	{'.', '.', '.', '2', '7', '.', '.', '1', '.'},
	//	{'3', '.', '.', '.', '5', '9', '.', '.', '.'},
	//	{'.', '9', '.', '8', '.', '.', '2', '.', '4'},
	//	{'6', '8', '5', '.', '.', '.', '.', '.', '9'},
	//	{'.', '.', '4', '.', '.', '.', '.', '.', '.'},
	//	{'.', '.', '.', '.', '.', '.', '.', '3', '.'},
	//	{'9', '.', '.', '.', '.', '.', '.', '.', '.'},
	//	{'.', '.', '1', '.', '9', '.', '7', '2', '.'},
	//	{'.', '.', '.', '5', '.', '.', '8', '.', '.'},
	//}
	boards := [][][]byte{
		{
			{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
			{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
			{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
			{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
			{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
			{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
			{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
			{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
			{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
		},
		{
			{'.', '.', '.', '2', '7', '.', '.', '1', '.'},
			{'3', '.', '.', '.', '5', '9', '.', '.', '.'},
			{'.', '9', '.', '8', '.', '.', '2', '.', '4'},
			{'6', '8', '5', '.', '.', '.', '.', '.', '9'},
			{'.', '.', '4', '.', '.', '.', '.', '.', '.'},
			{'.', '.', '.', '.', '.', '.', '.', '3', '.'},
			{'9', '.', '.', '.', '.', '.', '.', '.', '.'},
			{'.', '.', '1', '.', '9', '.', '7', '2', '.'},
			{'.', '.', '.', '5', '.', '.', '8', '.', '.'},
		},
		{{'.', '.', '9', '7', '4', '8', '.', '.', '.'}, {'7', '.', '.', '.', '.', '.', '.', '.', '.'}, {'.', '2', '.', '1', '.', '9', '.', '.', '.'}, {'.', '.', '7', '.', '.', '.', '2', '4', '.'}, {'.', '6', '4', '.', '1', '.', '5', '9', '.'}, {'.', '9', '8', '.', '.', '.', '3', '.', '.'}, {'.', '.', '.', '8', '.', '3', '.', '2', '.'}, {'.', '.', '.', '.', '.', '.', '.', '.', '6'}, {'.', '.', '.', '2', '7', '5', '9', '.', '.'}},
		{{'1', '.', '.', '.', '7', '.', '.', '3', '.'}, {'8', '3', '.', '6', '.', '.', '.', '.', '.'}, {'.', '.', '2', '9', '.', '.', '6', '.', '8'}, {'6', '.', '.', '.', '.', '4', '9', '.', '7'}, {'.', '9', '.', '.', '.', '.', '.', '5', '.'}, {'3', '.', '7', '5', '.', '.', '.', '.', '4'}, {'2', '.', '3', '.', '.', '9', '1', '.', '.'}, {'.', '.', '.', '.', '.', '2', '.', '4', '3'}, {'.', '4', '.', '.', '8', '.', '.', '.', '9'}},
	}
	for _, b := range boards {
		printBoard(b)
		solveSudoku(b)

		printBoard(b)
		println()
	}
}

func printBoard(b [][]byte) {
	for i, r := range b {
		if i == 3 || i == 6 {
			fmt.Println("-----------")
		}
		for j, c := range r {
			if j == 3 || j == 6 {
				fmt.Print("|")
			}
			fmt.Printf("%c", c)
		}
		fmt.Println()
	}
}

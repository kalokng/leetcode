package main

import "fmt"

func main() {
	fmt.Println(highestPeak([][]int{
		{0, 1},
		{0, 0},
	}))
	fmt.Println(highestPeak([][]int{
		{0, 0, 1},
		{1, 0, 0},
		{0, 0, 0},
	}))
}

func highestPeak(isWater [][]int) [][]int {
	type P struct {
		x int
		y int
	}
	m := map[P]struct{}{}
	for i := range isWater {
		for j := range isWater[i] {
			isWater[i][j]--
			if isWater[i][j] == 0 {
				m[P{i, j}] = struct{}{}
			}
		}
	}
	fill := func(nm map[P]struct{}, x, y, v int) {
		if isWater[x][y] >= 0 {
			return
		}
		isWater[x][y] = v
		nm[P{x, y}] = struct{}{}
	}
	for len(m) > 0 {
		nm := map[P]struct{}{}
		for p := range m {
			v := isWater[p.x][p.y] + 1
			if p.x > 0 {
				fill(nm, p.x-1, p.y, v)
			}
			if p.y > 0 {
				fill(nm, p.x, p.y-1, v)
			}
			if p.x < len(isWater)-1 {
				fill(nm, p.x+1, p.y, v)
			}
			if p.y < len(isWater[0])-1 {
				fill(nm, p.x, p.y+1, v)
			}
		}
		m = nm
	}
	return isWater
}

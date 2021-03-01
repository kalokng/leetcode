package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(getCollisionTimes([][]int{{1, 2}, {2, 1}, {4, 3}, {7, 2}}))
	fmt.Println(getCollisionTimes([][]int{{3, 4}, {5, 4}, {6, 3}, {9, 1}}))
}

type Stk struct {
	p, s int
	t    float64
}

func getCollisionTimes(cars [][]int) []float64 {
	t := make([]float64, len(cars))
	var stk []Stk
	for i := len(cars) - 1; i >= 0; i-- {
		p, s := cars[i][0], cars[i][1]
		var hit float64
		for len(stk) > 0 {
			j := len(stk) - 1
			if stk[j].s >= s {
				stk = stk[:j]
				continue
			}
			hit = float64(stk[j].p-p) / float64(s-stk[j].s)
			if hit < stk[j].t {
				break
			}
			stk = stk[:j]
		}
		if len(stk) == 0 {
			t[i] = -1
			stk = append(stk, Stk{p, s, math.Inf(1)})
			continue
		}
		t[i] = hit
		stk = append(stk, Stk{p, s, hit})
	}
	return t
}

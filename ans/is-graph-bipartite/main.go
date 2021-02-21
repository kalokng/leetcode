package main

import "fmt"

func main() {
	fmt.Println(isBipartite([][]int{
		{1, 2, 3},
		{0, 2},
		{0, 1, 3},
		{0, 2},
	}))
	fmt.Println(isBipartite([][]int{
		{1, 3},
		{0, 2},
		{1, 3},
		{0, 2},
	}))
	fmt.Println(isBipartite([][]int{
		{1},
		{0, 3},
		{3},
		{1, 2},
	}))
	fmt.Println(isBipartite([][]int{
		{4, 1}, {0, 2}, {1, 3}, {2, 4}, {3, 0},
	}))
}

func isBipartite(graph [][]int) bool {
	t := make([]bool, len(graph))
	for i, e := range graph {
		if t[i] {
			continue
		}
		a, b := map[int]bool{i: true}, map[int]bool{}
		l := map[int][]int{i: e}
		for len(l) > 0 {
			nl := map[int][]int{}
			for i, e := range l {
				t[i] = true
				for _, j := range e {
					if t[j] {
						continue
					}
					if a[j] {
						return false
					}
					b[j] = true
					nl[j] = graph[j]
				}
			}
			a, b = b, a
			l = nl
		}
	}
	return true
}

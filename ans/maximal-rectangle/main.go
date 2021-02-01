package main

import "fmt"

type pair struct {
	idx    int
	height int
}

func largestRectangleArea(heights []int) int {
	var stack []pair
	max := 0
	for i, h := range heights {
		s := i
		for len(stack) > 0 {
			v := stack[len(stack)-1]
			if v.height <= h {
				break
			}
			area := v.height * (i - v.idx)
			if max < area {
				max = area
			}
			s = v.idx
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, pair{s, h})
	}

	l := len(heights)
	for _, p := range stack {
		area := p.height * (l - p.idx)
		if area > max {
			max = area
		}
	}
	return max
}

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	height := make([]int, len(matrix[0]))
	max := 0
	for _, r := range matrix {
		for i, v := range r {
			if v == '0' {
				height[i] = 0
			} else {
				height[i]++
			}
		}
		area := largestRectangleArea(height)
		if area > max {
			max = area
		}
	}
	return max
}

func main() {
	fmt.Println(maximalRectangle([][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}))
}

package main

func main() {
}

func distributeCandies(candyType []int) int {
	n := len(candyType) / 2
	m := make(map[int]struct{})
	for _, v := range candyType {
		m[v] = struct{}{}
	}
	if n < len(m) {
		return n
	}
	return len(m)
}

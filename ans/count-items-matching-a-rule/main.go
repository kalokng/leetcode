package main

func main() {
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	n := 0
	var k int
	switch ruleKey {
	case "type":
		k = 0
	case "color":
		k = 1
	case "name":
		k = 2
	}
	for _, i := range items {
		if i[k] == ruleValue {
			n++
		}
	}
	return n
}

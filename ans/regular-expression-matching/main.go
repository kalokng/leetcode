package main

func isMatch(s string, p string) bool {
	m, _ := regexp("^"+p+"$", s)
	return m
}

func main() {
}

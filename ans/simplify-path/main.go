package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(simplifyPath("/home/"))
	fmt.Println(simplifyPath("/../"))
	fmt.Println(simplifyPath("/home//foo/"))
	fmt.Println(simplifyPath("/a/./b/../../c/"))
}

func simplifyPath(path string) string {
	var dirs []string
	ps := strings.Split(path, "/")
	for _, p := range ps {
		switch p {
		case "", ".":
		case "..":
			if len(dirs) > 0 {
				dirs = dirs[:len(dirs)-1]
			}
		default:
			dirs = append(dirs, p)
		}
	}
	return "/" + strings.Join(dirs, "/")
}

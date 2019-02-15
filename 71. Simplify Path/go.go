package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(simplifyPath("/home"))
	fmt.Println(simplifyPath("/../"))
	fmt.Println(simplifyPath("/home//foo/"))
	fmt.Println(simplifyPath("/a/../../b/../c//.//"))
	fmt.Println(simplifyPath("/a//b////c/d//././/.."))

}

func simplifyPath(path string) string {
	tmpPath := strings.Split(path, "/")
	spath := make([]string, 0)
	for _, v := range tmpPath {
		switch v {
		case "", ".":
		case "..":
			if len(spath) == 0 {
				break
			}
			spath = spath[0 : len(spath)-1]
		default:
			spath = append(spath, v)
		}
	}
	return "/" + strings.Join(spath, "/")
}

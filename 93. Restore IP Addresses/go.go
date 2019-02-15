package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(restoreIpAddresses("010010"))
	fmt.Println(restoreIpAddresses("25525511135"))
}

func restoreIpAddresses(s string) []string {
	res := make([]string, 0)

	for a := 1; a < 4; a++ {
		for b := 1; b < 4; b++ {
			for c := 1; c < 4; c++ {
				for d := 1; d < 4; d++ {
					if a+b+c+d == len(s) {
						ai, _ := strconv.Atoi(s[0:a])
						bi, _ := strconv.Atoi(s[a : a+b])
						ci, _ := strconv.Atoi(s[a+b : a+b+c])
						di, _ := strconv.Atoi(s[a+b+c : a+b+c+d])
						if ai < 256 && bi < 256 && ci < 256 && di < 256 && len(fmt.Sprintf("%d.%d.%d.%d", ai, bi, ci, di)) == len(s)+3 {
							res = append(res, fmt.Sprintf("%d.%d.%d.%d", ai, bi, ci, di))
						}
					}
				}
			}
		}
	}

	return res
}

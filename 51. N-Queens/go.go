package main

import (
	"fmt"
)

func main() {
	fmt.Println(solveNQueens(4))
}

func solveNQueens(n int) [][]string {
	res := make([][]string, 0)
	poss := make([]int, n)
	dfs(n, 0, 0, poss, &res)
	return res
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func checkPos(poss []int, row, pos int) bool {
	if row == 0 {
		return true
	}
	for i := 0; i < row; i++ {
		if poss[i] == pos || abs(row-i) == abs(pos-poss[i]) {
			return false
		}
	}
	return true
}

func dfs(n, row, pos int, poss []int, res *[][]string) {
	if row == n {
		*res = append(*res, genRes(poss))
		return
	}
	if pos == n {
		return
	}
	if checkPos(poss, row, pos) {
		poss[row] = pos
		dfs(n, row+1, 0, poss, res)
	}
	dfs(n, row, pos+1, poss, res)
}

func genRes(poss []int) []string {
	res := make([]string, len(poss))
	for i := range res {
		s := ""
		for j := 0; j < len(poss); j++ {
			if j == poss[i] {
				s += "Q"
			} else {
				s += "."
			}
		}
		res[i] = s
	}

	return res
}

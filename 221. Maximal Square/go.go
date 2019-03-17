package main

import "fmt"

func main() {
	matrix := make([][]byte, 4)
	matrix[0] = []byte{'1', '0', '1', '0', '0'}
	matrix[1] = []byte{'1', '0', '1', '1', '1'}
	matrix[2] = []byte{'1', '1', '1', '1', '1'}
	matrix[3] = []byte{'1', '0', '0', '1', '0'}
	fmt.Println(maximalSquare(matrix))
}

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	rows, cows := len(matrix), len(matrix[0])
	dp := make([][]int, rows)
	for i := 0; i < rows; i++ {
		dp[i] = make([]int, cows)
	}
	var res int
	for i := 0; i < rows; i++ {
		for j := 0; j < cows; j++ {
			if matrix[i][j] == '1' {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = minNum(dp[i][j-1], dp[i-1][j], dp[i-1][j-1]) + 1
				}
				if dp[i][j] > res {
					res = dp[i][j]
				}
			}
		}
	}

	return res*res
}

func minNum(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}
	res := 0xfffff
	for _, v := range nums {
		if v < res {
			res = v
		}
	}
	return res
}

/*
 * @lc app=leetcode.cn id=221 lang=golang
 *
 * [221] 最大正方形
 */

// @lc code=start
func maximalSquare(matrix [][]byte) int {
	lengthMatrix := make([][]int, len(matrix))
	for i := range lengthMatrix {
		lengthMatrix[i] = make([]int, len(matrix[0]))
	}

	var res int
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] != '1' {
				continue
			}

			if i == 0 || j == 0 {
				lengthMatrix[i][j] = 1
			} else {
				lengthMatrix[i][j] = min(lengthMatrix[i-1][j], lengthMatrix[i][j-1], lengthMatrix[i-1][j-1]) + 1
			}

			if lengthMatrix[i][j] > res {
				res = lengthMatrix[i][j]
			}
		}
	}

	return res * res
}

func min(ints ...int) int {
	minVal := ints[0]
	for _, i := range ints {
		if i < minVal {
			minVal = i
		}
	}

	return minVal
}

// @lc code=end


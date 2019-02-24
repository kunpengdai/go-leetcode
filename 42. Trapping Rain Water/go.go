package main

import "fmt"

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}
/*
* 对于每一个位置，如果此位置左边的最大值与右边的最大值，都大于自己本身，那么该位置能存储水
* 则该问题求出每个位置，左边的最大值和右边的最大值即可
*/
func trap(height []int) int {
	maxLeft, maxRight, length, res := make([]int, len(height)), make([]int, len(height)), len(height), 0
	if length < 3 {
		return 0
	}
	maxLeft[0] = height[0]
	maxRight[length-1] = height[length-1]
	for i := 1; i < length; i++ {
		maxLeft[i] = max(maxLeft[i-1], height[i])
		maxRight[length-1-i] = max(maxRight[length-i], height[length-i-1])
	}
	for i := range height {
		if height[i] < maxLeft[i] && height[i] < maxRight[i] {
			res += min(maxLeft[i], maxRight[i]) - height[i]
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

package main

import "fmt"

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res := nums[0]
	tmpRes := 0
	for _, num := range nums {
		tmpRes += num
		if tmpRes > res {
			res = tmpRes
		}
		if tmpRes < 0 {
			tmpRes = 0
		}
	}
	return res
}

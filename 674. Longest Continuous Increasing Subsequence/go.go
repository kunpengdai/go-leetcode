package main

import "fmt"

func main() {
	fmt.Println(findLengthOfLCIS([]int{1, 2, 3, 3, 5, 4, 3, 0, 5, 7, 8, 9}))
}

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxLength := 1
	tmpLength := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			if tmpLength > maxLength {
				maxLength = tmpLength
			}
			tmpLength = 1
		} else {
			tmpLength++
		}
	}
	if tmpLength > maxLength {
		maxLength = tmpLength
	}
	return maxLength
}

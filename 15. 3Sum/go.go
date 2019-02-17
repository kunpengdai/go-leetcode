package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 0, -1, 3, -1, 5}
	// threeSum(nums)
	fmt.Println(threeSum(nums))
}

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	if len(nums) < 3 {
		return res
	}

	sort.Ints(nums)

	for i := 0; i < len(nums)-2;i=findNext(nums,i) {
		start := i + 1
		end := len(nums) - 1
		for start < end {
			sum := nums[i] + nums[start] + nums[end]
			if sum == 0 {
				item := make([]int, 3)
				item[0] = nums[i]
				item[1] = nums[start]
				item[2] = nums[end]
				start = findNext(nums, start)
				end = findPre(nums, end)
				res = append(res, item)
			} else if sum < 0 {
				start = findNext(nums, start)
			} else {
				end = findPre(nums, end)
			}
		}
	}

	return res
}

func findNext(nums []int, index int) int {
	for i := index + 1; i < len(nums); i++ {
		if nums[i] != nums[index] {
			return i
		}
		fmt.Println("")
	}
	return len(nums)
}

func findPre(nums []int, index int) int {
	for i := index - 1; i >= 0; i-- {
		if nums[i] != nums[index] {
			return i
		}
	}
	return -1
}

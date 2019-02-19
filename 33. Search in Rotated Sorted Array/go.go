package main

import "fmt"

func main() {
	fmt.Println(search([]int{3,1}, 1))
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	firstNum := nums[0]
	if firstNum == target {
		return 0
	}
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := (low + high) >> 1
		switch {
		case nums[mid] == target:
			return mid
		case nums[mid] >= nums[low]:
			if target >= nums[low] && target <= nums[mid] {
				high = mid-1
			} else {
				low = mid+1
			}
		case nums[mid] < nums[low]:
			if target >= nums[mid] && target <= nums[high] {
				low = mid+1
			} else {
				high = mid-1
			}
		}
	}
	return -1
}

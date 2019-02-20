package main

import "fmt"

func main() {
	fmt.Println(longestConsecutive([]int{1, 5, 3, 76, 2, 8, 4, 9, 6, 11, 23, 12, 13, 15, 14, 16, 18, 17, 19}))
}

func longestConsecutive(nums []int) int {
	existed := make(map[int]bool, len(nums))
	for _, num := range nums {
		existed[num] = true
	}
	res := 0
	for _, num := range nums {
		tmpRes := 1
		if existed[num] {
			existed[num] = false
			for i := num + 1; existed[i]; i++ {
				existed[i] = false
				tmpRes++
			}
			for i := num - 1; existed[i]; i-- {
				existed[i] = false
				tmpRes++
			}
		}
		if tmpRes > res {
			res = tmpRes
		}
	}
	return res
}

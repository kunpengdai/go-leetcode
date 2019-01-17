package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(triangleNumber([]int{1, 2, 3, 4, 5, 6}))
}

/*解题思路：
先排序，从前往后先选定两个边，再选数字小于此两边之和的边；时间复杂度O（n^2*logn）
*/
func triangleNumber(nums []int) int {
	res := 0
	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	nums = append(nums, 0xffffffff)
	for i := 0; i < len(nums)-3; i++ {
		for j := i + 1; j < len(nums)-2; j++ {
			id := sort.Search(len(nums), func(ix int) bool {
				return nums[ix] >= (nums[i] + nums[j])
			}) - 1
			fmt.Println(i, j, id)
			if id > j {
				res = res + id - j
			}
		}
	}
	return res
}

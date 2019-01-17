package main

import (
	"fmt"
	"strconv"
)

func main() {
	nums := []int{0, 1, 2, 4, 5, 7}
	fmt.Println(summaryRanges(nums))
}

type pair struct {
	start int
	end   int
}

func summaryRanges(nums []int) []string {
	res := make([]string, 0)
	pairs := make([]pair, 0)
	if len(nums) == 0 {
		return res
	}
	for index, num := range nums {
		if index > 0 && num == nums[index-1]+1 {
			pairs[len(pairs)-1].end = num
		} else {
			pairs = append(pairs, pair{start: num, end: num})
		}
	}
	for _, p := range pairs {
		if p.start == p.end {
			res = append(res, strconv.Itoa(p.start))
		} else {
			res = append(res, fmt.Sprintf("%d->%d", p.start, p.end))
		}
	}
	return res
}

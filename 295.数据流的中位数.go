/*
 * @lc app=leetcode.cn id=295 lang=golang
 *
 * [295] 数据流的中位数
 */

// @lc code=start
type MedianFinder struct {
	nums []int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{nums: make([]int, 0)}
}

func (this *MedianFinder) AddNum(num int) {

	low, high := 0, len(this.nums)-1
	for low <= high {
		mid := (low + high) / 2
		if this.nums[mid] < num {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if low == len(this.nums) {
		this.nums = append(this.nums, num)
		return
	}

	this.nums = append(this.nums[:low], append([]int{num}, this.nums[low:]...)...)
}

func (this *MedianFinder) FindMedian() float64 {
	if len(this.nums) == 0 {
		return 0.0
	}

	if len(this.nums)&1 == 1 {
		return float64(this.nums[len(this.nums)/2])
	} else {
		return (float64(this.nums[len(this.nums)/2-1]) + float64(this.nums[len(this.nums)/2])) / 2
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
// @lc code=end


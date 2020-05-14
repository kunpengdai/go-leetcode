/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 */

// @lc code=start
func maxSlidingWindow(nums []int, k int) []int {

	if len(nums) == 0 || k > len(nums) {
		return []int{}
	}
	dq := make(dqueue, 0)
	dq.pushback(nums[0])
	var maxPos int
	for i := 1; i < k; i++ {
		if nums[i] > nums[maxPos] {
			maxPos = i
			dq = make(dqueue, 1)
			dq[0] = nums[i]
		} else {
			dq.pushback(nums[i])
		}
	}
	res := make([]int, 1)
	res[0] = dq.getMax()

	for i := k; i < len(nums); i++ {
		if nums[i] > nums[maxPos] {
			dq = make(dqueue, 1)
			dq[0] = nums[i]
			maxPos = i
		} else if i > maxPos+k-1 {
			dq.popfront()
			dq.pushback(nums[i])
			maxPos+=(1+dq.remainMax())
		} else {
			dq.pushback(nums[i])
		}
		res = append(res, dq.getMax())
	}

	return res
}

type dqueue []int

func (d *dqueue) pushback(i int) {
	*d = append(*d, i)
}

func (d *dqueue) popfront() {
	*d = (*d)[1:]
}

func (d *dqueue) remainMax()int {
	var maxPos int
	for i := 1; i < len(*d); i++ {
		if (*d)[maxPos] < (*d)[i] {
			maxPos = i
		}
	}
	*d = (*d)[maxPos:]
	return maxPos
}

func (d *dqueue) getMax() int{
	return (*d)[0]
}

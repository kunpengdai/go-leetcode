/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 */

// @lc code=start
type numSlice []int

func (n numSlice) Len() int { return len(n) }

func (n numSlice) Less(i, j int) bool { return n[i] > n[j] }

func (n *numSlice) Swap(i, j int) { (*n)[i], (*n)[j] = (*n)[j], (*n)[i] }

func (n *numSlice) Push(i interface{}) {
	*n = append(*n, i.(int))
}

func (n *numSlice) Pop() interface{} {
	length := n.Len()
	res := (*n)[length-1]
	*n = (*n)[:length-1]
	return res
}

func (n numSlice) Top() int {
	return n[0]
}

func findKthLargest(nums []int, k int) int {

	ns := numSlice(nums)
	heap.Init(&ns)

	var res int
	for k > 0 {
		res = heap.Pop(&ns).(int)
		k--
	}

	return res
}


// @lc code=end


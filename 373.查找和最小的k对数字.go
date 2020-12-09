/*
 * @lc app=leetcode.cn id=373 lang=golang
 *
 * [373] 查找和最小的K对数字
 */

// @lc code=start
type ij struct {
	s1    *[]int
	s2    *[]int
	s1idx int
	s2idx int
}

type ijSlice []ij

func (is ijSlice) Less(i, j int) bool {
	return (*is[i].s1)[is[i].s1idx]+(*is[i].s2)[is[i].s2idx] < (*is[j].s1)[is[j].s1idx]+(*is[j].s2)[is[j].s2idx]
}

func (is ijSlice) Len() int {
	return len(is)
}

func (is ijSlice) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

func (is *ijSlice) Push(x interface{}) {
	*is = append(*is, x.(ij))
}

func (is *ijSlice) Pop() interface{} {
	length := len(*is)
	x := (*is)[length-1]
	*is = (*is)[:length-1]
	return x
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {

	if len(nums1)==0 || len(nums2)==0 {
		return nil
	}

	res := make([][]int, 0, k)

	ijs := ijSlice{}
	heap.Init(&ijs)

	for i := 0; i < len(nums1); i++ {
		heap.Push(&ijs, ij{
			s1:    &nums1,
			s2:    &nums2,
			s1idx: i,
			s2idx: 0,
		})
	}

	for kp := 0; kp < k; kp++ {

		if ijs.Len() == 0 {
			break
		}

		minV := heap.Pop(&ijs).(ij)
		res = append(res, []int{nums1[minV.s1idx], nums2[minV.s2idx]})
		if minV.s2idx < len(nums2)-1 {
			heap.Push(&ijs, ij{
				s1:    &nums1,
				s2:    &nums2,
				s1idx: minV.s1idx,
				s2idx: minV.s2idx + 1,
			})
		}
	}

	return res
}

// @lc code=end


/*
 * @lc app=leetcode.cn id=480 lang=golang
 *
 * [480] 滑动窗口中位数
 */

// @lc code=start
func medianSlidingWindow(nums []int, k int) []float64 {
	kSlice := IntSlice{
		data: make([]int, k),
	}
	for i := 0; i < k; i++ {
		kSlice.data[i] = nums[i]
	}
	kSlice.sort()

	res := make([]float64, len(nums)-k+1)
	res[0] = kSlice.getMiddle()
	for i := k; i < len(nums); i++ {
		kSlice.deleteData(nums[i-k])
		kSlice.insertData(nums[i])
		res[i-k+1] = kSlice.getMiddle()
	}

	return res
}

type IntSlice struct {
	data []int
}

func (s *IntSlice) sort() {
	sort.Slice(s.data, func(i, j int) bool {
		return s.data[i] < s.data[j]
	})
}

func (s IntSlice) getMiddle() float64 {
	length := len(s.data)
	if length == 0 {
		return 0
	}
	if length&1 == 1 {
		return float64(s.data[length>>1])
	}

	return float64(s.data[length>>1]+s.data[(length>>1)-1]) / 2
}

func (s *IntSlice) insertData(num int) {

	if len(s.data)==0{
		s.data = []int{num}
		return
	}

	l, h := 0, len(s.data)-1
	for l <= h {
		mid := (l + h) >> 1
		if s.data[mid] <= num {
			if mid == len(s.data)-1 || s.data[mid+1] >= num {
				newS := make([]int, len(s.data)-mid-1)
				copy(newS, s.data[mid+1:])
				s.data = append(s.data[0:mid+1], num)
				s.data = append(s.data, newS...)
				return
			}
			l = mid + 1
		} else if s.data[mid] >= num {
			if mid == 0 || s.data[mid-1] <= num {
				newS := make([]int, len(s.data)-mid)
				copy(newS, s.data[mid:])
				s.data = append(s.data[0:mid], num)
				s.data = append(s.data, newS...)
				return
			}
			h = mid - 1
		}
	}
	return
}

func (s *IntSlice) deleteData(num int) {
	l, h := 0, len(s.data)-1
	for l <= h {
		mid := (l + h) >> 1
		if s.data[mid] == num {
			s.data = append(s.data[0:mid], s.data[mid+1:]...)
			return
		} else if s.data[mid] < num {
			l = mid + 1
		} else {
			h = mid - 1
		}
	}
}
// @lc code=end


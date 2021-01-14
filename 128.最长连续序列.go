/*
 * @lc app=leetcode.cn id=128 lang=golang
 *
 * [128] 最长连续序列
 */

// @lc code=start
type fathers []int

func longestConsecutive(nums []int) int {

	m := make(map[int]int, len(nums))
	fs := fathers(make([]int, len(nums)))
	for i := 0; i < len(fs); i++ {
		fs[i] = i
	}

	for i := 0; i < len(nums); i++ {

		// remove existing num
		if _, ok := m[nums[i]]; ok {
			continue
		}

		m[nums[i]] = i

		for _, neighbor := range []int{
			nums[i] - 1, nums[i] + 1,
		} {
			if pos, ok := m[neighbor]; ok {
				fs.merge(i, pos)
			}
		}
	}

	var res int
	count := make(map[int]int, len(nums))
	for i := range fs {
		count[fs.find(i)]++
		if count[fs.find(i)] > res {
			res = count[fs.find(i)]
		}
	}

	return res
}

func (f fathers) find(i int) int {
	if f[i] == i {
		return i
	}

	return f.find(f[i])
}

func (f *fathers) merge(i, j int) {

	fi := f.find(i)
	fj := f.find(j)
	if fi == fj {
		return
	}

	(*f)[fi] = fj
}

// @lc code=end


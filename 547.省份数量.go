/*
 * @lc app=leetcode.cn id=547 lang=golang
 *
 * [547] 省份数量
 */

// @lc code=start

type fathers []int

func findCircleNum(isConnected [][]int) (ans int) {

	fs := fathers(make([]int, len(isConnected)))
	for i := 0; i < len(fs); i++ {
		fs[i] = i
	}

	for i, s := range isConnected {
		for j, c := range s {
			if c == 1 {
				fs.merge(i, j)
			}
		}
	}

	for i, v := range fs {
		if i == v {
			ans++
		}
	}

	return
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


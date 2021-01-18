/*
 * @lc app=leetcode.cn id=871 lang=golang
 *
 * [871] 最低加油次数
 */

// @lc code=start

type fuels []int

func (f fuels) Len() int { return len(f) }

func (f fuels) Less(i, j int) bool {
	return f[i] > f[j]
}

func (f *fuels) Swap(i, j int) {
	(*f)[i], (*f)[j] = (*f)[j], (*f)[i]
}

func (f *fuels) Pop() interface{} {
	length := f.Len()
	var res int
	res = (*f)[length-1]
	*f = (*f)[:length-1]
	return res
}

func (f *fuels) Push(x interface{}) {
	*f = append(*f, x.(int))
}

func minRefuelStops(target int, startFuel int, stations [][]int) int {
	fs := fuels(make([]int, 0, len(stations)))
	heap.Init(&fs)

	heap.Push(&fs, startFuel)
	var farthestPos, stationIndex, res int
	for len(fs) != 0 {
		maxFuel := heap.Pop(&fs).(int)
		farthestPos += maxFuel
		if farthestPos >= target {
			return res
		}

		res++
		for stationIndex < len(stations) {
			if stations[stationIndex][0] <= farthestPos {
				heap.Push(&fs, stations[stationIndex][1])
				stationIndex++
			} else {
				break
			}
		}
	}

	return -1
}

// @lc code=end


/*
 * @lc app=leetcode.cn id=313 lang=golang
 *
 * [313] 超级丑数
 */

// @lc code=start

type uglyNums []int

func (s uglyNums) Len() int           { return len(s) }
func (s uglyNums) Less(i, j int) bool { return s[i] < s[j] }
func (s *uglyNums) Swap(i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}
func (s *uglyNums) Push(x interface{}) {
	(*s) = append(*s, x.(int))
}
func (s *uglyNums) Pop() interface{} {
	length := s.Len()
	res := (*s)[length-1]
	*s = (*s)[:length-1]
	return res
}

func nthSuperUglyNumber(n int, primes []int) int {
	us := uglyNums(make([]int, 0, n))
	heap.Push(&us, 1)
	numMap := make(map[int]bool, n)

	var min int
	for n > 0 {
		n--
		min = heap.Pop(&us).(int)
		for _, prime := range primes {
			if !numMap[min*prime] {
				numMap[min*prime] = true
				heap.Push(&us, min*prime)
			}
		}
	}

	return min
}

// @lc code=end


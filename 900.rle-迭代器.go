/*
 * @lc app=leetcode.cn id=900 lang=golang
 *
 * [900] RLE 迭代器
 */

// @lc code=start
type RLEIterator struct {
	index     int
	rle       []int
	remainNum int
}

func Constructor(A []int) RLEIterator {
	if len(A) == 0 {
		return RLEIterator{}
	}

	return RLEIterator{
		index:     1,
		rle:       A,
		remainNum: A[0],
	}
}

func (this *RLEIterator) Next(n int) int {

	if this.remainNum >= n {
		this.remainNum -= n
		return this.rle[this.index]
	}

	n -= this.remainNum
	this.remainNum = 0
	if this.index == len(this.rle)-1 {
		return -1
	}

	this.remainNum = this.rle[this.index+1]
	this.index += 2

	return this.Next(n)
}

/**
 * Your RLEIterator object will be instantiated and called as such:
 * obj := Constructor(A);
 * param_1 := obj.Next(n);
 */
// @lc code=end


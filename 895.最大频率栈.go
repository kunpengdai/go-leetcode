/*
 * @lc app=leetcode.cn id=895 lang=golang
 *
 * [895] 最大频率栈
 */

// @lc code=start
type FreqStack struct {
	maxFreq  map[int]int //map[num]frep
	freqList map[int][]int
	maxF     int
}

func Constructor() FreqStack {
	return FreqStack{
		maxFreq:  make(map[int]int),
		freqList: make(map[int][]int),
		maxF:     0,
	}
}

func (this *FreqStack) Push(x int) {
	this.maxFreq[x]++
	this.freqList[this.maxFreq[x]] = append(this.freqList[this.maxFreq[x]], x)
	if this.maxF < this.maxFreq[x] {
		this.maxF = this.maxFreq[x]
	}
}

func (this *FreqStack) Pop() int {
	maxList := this.freqList[this.maxF]
	res := maxList[len(maxList)-1]
	if len(maxList) == 1 {
		delete(this.freqList, this.maxF)
		this.maxF--
	} else {
		this.freqList[this.maxF] = this.freqList[this.maxF][0 : len(this.freqList[this.maxF])-1]
	}
	if freq, _ := this.maxFreq[res]; freq == 1 {
		delete(this.maxFreq, res)
	} else {
		this.maxFreq[res]--
	}

	return res
}

/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 */
// @lc code=end


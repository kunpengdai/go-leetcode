/*
 * @lc app=leetcode.cn id=502 lang=golang
 *
 * [502] IPO
 */

// @lc code=start
type capitalProfitSlice []capProfit
type profitSlice []int

func (c capitalProfitSlice) Len() int {
	return len(c)
}

func (p profitSlice) Len() int {
	return len(p)
}

func (c capitalProfitSlice) Less(i, j int) bool {
	return c[i].capital < c[j].capital
}

func (p profitSlice) Less(i, j int) bool {
	return p[i] > p[j]
}

func (c capitalProfitSlice) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (p profitSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (c *capitalProfitSlice) Push(item interface{}) {
	*c = append(*c, item.(capProfit))
}

func (p *profitSlice) Push(item interface{}) {
	*p = append(*p, item.(int))
}

func (c *capitalProfitSlice) Pop() interface{} {
	length := len(*c)
	res := (*c)[length-1]
	*c = (*c)[:length-1]
	return res
}

func (c *capitalProfitSlice) Top() capProfit {
	return (*c)[0]
}

func (p *profitSlice) Pop() interface{} {
	length := len(*p)
	res := (*p)[length-1]
	*p = (*p)[:length-1]
	return res
}

type capProfit struct {
	capital int
	profit  int
}

func findMaximizedCapital(k int, W int, Profits []int, Capital []int) int {

	pcs := capitalProfitSlice(make([]capProfit, 0, len(Profits)))
	ps := profitSlice(make([]int, 0, k))

	for i := 0; i < len(Profits); i++ {
		heap.Push(&pcs, capProfit{
			capital: Capital[i],
			profit:  Profits[i],
		})
	}

	for k > 0 {

		for len(pcs) != 0 && pcs.Top().capital <= W {
			minCapital := heap.Pop(&pcs).(capProfit)
			heap.Push(&ps, minCapital.profit)
		}

		if len(ps) == 0 {
			break
		}

		maxProfit := heap.Pop(&ps).(int)
		W += maxProfit
		k--
	}

	return W
}

// @lc code=end

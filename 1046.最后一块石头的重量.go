/*
 * @lc app=leetcode.cn id=1046 lang=golang
 *
 * [1046] 最后一块石头的重量
 */

// @lc code=start
func lastStoneWeight(stones []int) int {
	for len(stones)>1{
		sort.Slice(stones,func (i,j int)bool  {
			return stones[i]>stones[j]	
		})
		if stones[0]>stones[1]{
			stones = append(stones,stones[0]-stones[1])
		}
		stones = stones[2:]
	}
	if len(stones)==1{
		return stones[0]
	}
	return 0
}
// @lc code=end


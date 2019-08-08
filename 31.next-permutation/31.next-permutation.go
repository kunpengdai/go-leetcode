/*
 * @lc app=leetcode id=31 lang=golang
 *
 * [31] Next Permutation
 *
 * https://leetcode.com/problems/next-permutation/description/
 *
 * algorithms
 * Medium (30.82%)
 * Likes:    2004
 * Dislikes: 638
 * Total Accepted:    256.9K
 * Total Submissions: 832.7K
 * Testcase Example:  '[1,2,3]'
 *
 * Implement next permutation, which rearranges numbers into the
 * lexicographically next greater permutation of numbers.
 * 
 * If such arrangement is not possible, it must rearrange it as the lowest
 * possible order (ie, sorted in ascending order).
 * 
 * The replacement must be in-place and use only constant extra memory.
 * 
 * Here are some examples. Inputs are in the left-hand column and its
 * corresponding outputs are in the right-hand column.
 * 
 * 1,2,3 → 1,3,2
 * 3,2,1 → 1,2,3
 * 1,1,5 → 1,5,1
 * 
 */
func nextPermutation(nums []int)  {
	pos,rep := -1,0
    for i:=len(nums)-2;i >= 0;i--{
		if nums[i]< nums[i+1]{
			pos = i
			break
		}
	}
	if pos == -1{
		sort.Ints(nums)
		return
	}
	for i:=len(nums)-1;i>pos;i--{
		if nums[i] >nums[pos]{
			rep = i
			break
		}
	}
	nums[pos],nums[rep] = nums[rep],nums[pos]
	sort.Ints(nums[pos+1:])
}


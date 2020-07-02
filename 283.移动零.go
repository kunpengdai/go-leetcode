/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start
func moveZeroes(nums []int)  {
	zp,np := 0,0
	for ;zp<len(nums);zp++{
		if nums[zp]==0{
			break
		}
	}

	for np=zp+1 ;np<len(nums);np++{
		if nums[np]!=0{
			break
		}
	}

	for zp<len(nums) && np < len(nums){
		nums[zp],nums[np] = nums[np],nums[zp]
		zp++
		for np=np+1;np<len(nums);np++{
			if nums[np]!=0{
				break
			}
		}
	}
}
// @lc code=end


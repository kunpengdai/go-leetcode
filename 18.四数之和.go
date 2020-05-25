/*
 * @lc app=leetcode.cn id=18 lang=golang
 *
 * [18] 四数之和
 */

// @lc code=start
func fourSum(nums []int, target int) [][]int {
	res := make([][]int, 0)
	if len(nums)<4{
		return res
	}

	sort.Slice(nums,func (i,j int)bool  {
		return nums[i]<nums[j]	
	})

	for i:=0;i<len(nums)-2;i++{
		for j := i+2;j<len(nums);j++{
			num := nums[i]+nums[j]
			k,t:= i+1,j-1
			for k<t{
				tmp := num + nums[k] + nums[t]
				if  tmp== target{
					app(&res,[]int{nums[i],nums[k],nums[t],nums[j]})
					k++
					t--
				}else if tmp < target{
					k++
				}else{
					t--
				}
			}
		}
	} 

	return res
}

func app(res *[][]int,ns []int){
	for i:= len(*res)-1;i>=0;i--{
		if ns[0] == (*res)[i][0] &&
		ns[1] == (*res)[i][1] &&
		ns[2] == (*res)[i][2] &&
		ns[3] == (*res)[i][3]{
			return
		}

		if ns[0]!=(*res)[i][0]{
			break
		}
	}

	*res = append(*res,ns)
}

// @lc code=end


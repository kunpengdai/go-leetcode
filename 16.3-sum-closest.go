/*
 * @lc app=leetcode id=16 lang=golang
 *
 * [16] 3Sum Closest
 */
func threeSumClosest(nums []int, target int) int {
	res := nums[0]+nums[1]+nums[2]
	sort.Ints(nums)
	for i:=0;i<len(nums)-2;i++{
		l,r := i+1,len(nums)-1
		for l<r{
			tmp := nums[i]+nums[l]+nums[r]
			if tmp == target {
				return tmp
			}
			res = closerNum(res,tmp,target)
			if tmp > target{
				r--
			}else{
				l++
			}
		}
	}
	return res	
}

func closerNum(a ,b,target int)int{
	if abs(a-target)<abs(b-target){
		return a
	}
	return b
}
func abs(a int)int{
	if a<0{
		return -a
	}
	return a
}

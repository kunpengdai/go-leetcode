/*
 * @lc app=leetcode id=238 lang=golang
 *
 * [238] Product of Array Except Self
 */
 func productExceptSelf(nums []int) []int {
    res := make([]int,len(nums))
    for i := range nums{
        res[i] = 1
    }
    for i:=1;i<len(nums);i++{
        res[i] = res[i-1]*nums[i-1]
    }
    tmp := nums[len(nums)-1]
    for i:=len(nums)-2;i>=0;i--{
        res[i]*=tmp
        tmp*=nums[i]
    }
    return res
}


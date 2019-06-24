package main

import "fmt"

func main(){
	fmt.Println(findMedianSortedArrays([]int{1,2,5},[]int{3,4,6,7}))
	fmt.Println(findMedianSortedArrays([]int{},[]int{3,4,6,7}))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	allLen := len(nums1)+len(nums2)
	allNum := make([]int, allLen)
	var i1,i2 int
	for i1<len(nums1) && i2<len(nums2){
		if nums1[i1]<nums2[i2]{
			allNum[i1+i2] = nums1[i1]
			i1++
		}else{
			allNum[i1+i2] = nums2[i2]
			i2++
		}
	}

	for i1<len(nums1){
		allNum[i1+i2] = nums1[i1]
		i1++
	}
	for i2<len(nums2){
		allNum[i1+i2] = nums2[i2]
		i2++
	}

	if len(allNum)&1==1{
		return float64(allNum[allLen/2])
	}
    return (float64(allNum[(allLen-1)/2])+float64(allNum[allLen/2]))/2
}
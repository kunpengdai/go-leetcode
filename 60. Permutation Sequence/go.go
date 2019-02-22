package main

import "fmt"

func main() {
	fmt.Println(getPermutation(3, 6))
}

func getPermutation(n int, k int) string {
	per := make([]int, 0)
	used := make([]bool, n)
	kth := 0
	dfs(k, &per, used, &kth)
	return genRes(per)
}

func dfs(k int, per *[]int, used []bool, kth *int) {
	if len(*per) == len(used) {
		*kth++
	}
	if k == *kth {
		used[len(used)-1]=false
		return
	}
	for index, u := range used {
		if u {
			continue
		}
		*per = append(*per, index+1)
		used[index] = true
		dfs(k, per, used, kth)
		if k==*kth{
			return
		}
		*per = (*per)[:len(*per)-1]
		used[index] = false
	}
}

func genRes(per []int) string {
	res := ""
	for _, v := range per {
		res += fmt.Sprintf("%d", v)
	}
	return res
}

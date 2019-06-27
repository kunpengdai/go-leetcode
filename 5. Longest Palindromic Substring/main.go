package main

import "fmt"

func main(){
	// fmt.Println(longestPalindrome("abc"))
	// fmt.Println(longestPalindrome("abcdc"))
	fmt.Println(longestPalindrome("bb"))
	fmt.Println(longestPalindrome("abcba"))
}

func longestPalindrome(s string) string {
	if s== ""{
		return ""
	}
	n := len(s)
	dp := make([][]bool, n)
	for d := range dp {
		dp[d] = make([]bool,n)
	}

	for i:=0;i<n;i++{
		dp[i][i] = true
	}

	for step:=1;step <= n-1;step++{
		for from:=0;from < n-step;from++{
			to := from+step
			if step==1{
				if s[to]==s[from]{
					dp[from][to] = true
				}
			}else{
				if s[to]==s[from] && dp[from+1][to-1]{
					dp[from][to] = true
				}
			}
		}
	}
	var maxLen,maxFrom,maxTo int
	for i := 0; i < n; i++ {
		for j:=i;j<n;j++{
			if dp[i][j] && j-i > maxLen{
				maxLen = j-i
				maxFrom = i
				maxTo = j
			}
		}
	}

	return s[maxFrom:maxTo+1]
}
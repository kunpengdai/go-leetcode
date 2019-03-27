package main

import (
	"fmt"
	"sort"
)

func main() {
	// envlops := [][]int{{4, 5}, {4, 6}, {6, 7}, {2, 3}, {1, 1}}
	// fmt.Println(maxEnvelopes(envlops))
	envlops1 := [][]int{{46, 89}, {50, 53}, {52, 68}, {72, 45}, {77, 81}}
	fmt.Println(maxEnvelopes(envlops1))
}

func maxEnvelopes(envelopes [][]int) int {
	res := 0
	if len(envelopes) == 0 {
		return 0
	}
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] < envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})
	fmt.Println(envelopes)
	dp := make([]int, len(envelopes))
	for i := range dp {
		dp[i] = 1
	}
	for i := 1; i < len(envelopes); i++ {
		for j := 0; j < i; j++ {
			if envelopes[i][1] > envelopes[j][1] && envelopes[i][0] > envelopes[j][0] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	for _, v := range dp {
		res = max(res, v)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

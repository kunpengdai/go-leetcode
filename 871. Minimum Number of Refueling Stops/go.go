package main

import "fmt"

func main() {
	station := [][]int{{10, 60}, {20, 30}, {30, 30}, {60, 40}}
	fmt.Println(minRefuelStops(100, 10, station))
}

/*minRefuelStops 解题思路
使用动态规划，dp[i]存储加了n次油能达到的最远的地方pos，dp[i+1]为pos内量最大的未加油的加油站 加油量f+pos，
直到所有的加油站都加油前，行路pos超出target即可。
*/
func minRefuelStops(target int, startFuel int, stations [][]int) int {
	dp := make([]int, len(stations)+1)
	used := make(map[int]bool, len(stations))
	if startFuel >= target {
		return 0
	}
	dp[0] = startFuel
	for i := 0; i < len(stations); i++ {
		lastPos := dp[i]
		maxFuel := 0
		maxFuelPos := 0
		for j := 0; j < len(stations); j++ {
			if lastPos >= stations[j][0] && !used[j] && maxFuel < stations[j][1] {
				maxFuel = stations[j][1]
				maxFuelPos = j
			}
		}
		dp[i+1] = lastPos + maxFuel
		used[maxFuelPos] = true
		if dp[i+1] >= target {
			return i + 1
		}
	}

	return -1
}

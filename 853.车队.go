/*
 * @lc app=leetcode.cn id=853 lang=golang
 *
 * [853] 车队
 */

// @lc code=start
type posSpeed struct {
	pos   int
	speed int
}

func carFleet(target int, position []int, speed []int) int {

	if len(position) <= 1 {
		return len(position)
	}

	posSpeeds := make([]posSpeed, len(position))
	for i := 0; i < len(position); i++ {
		posSpeeds[i] = posSpeed{
			pos:   position[i],
			speed: speed[i],
		}
	}

	sort.Slice(posSpeeds, func(i, j int) bool {
		return posSpeeds[i].pos > posSpeeds[j].pos
	})

	times := make([]float64, len(posSpeeds))
	times[0] = float64(target-posSpeeds[0].pos) / float64(posSpeeds[0].speed)
	for i := 1; i < len(posSpeeds); i++ {
		timeUsed := float64(target-posSpeeds[i].pos) / float64(posSpeeds[i].speed)
		if timeUsed < times[i-1] {
			times[i] = times[i-1]
		} else {
			times[i] = timeUsed
		}
	}

	fleets := 1
	for i := 1; i < len(times); i++ {
		if !equalFloat(times[i], times[i-1]) {
			fleets++
		}
	}

	return fleets
}

func equalFloat(a, b float64) bool {
	var diff float64
	if a > b {
		diff = a - b
	} else {
		diff = b - a
	}

	return diff < 0.000001
}

// @lc code=end


package main

import (
	"fmt"
	"sort"
)

func main() {
	position := []int{10, 8, 0, 5, 3}
	speed := []int{2, 4, 1, 1, 3}
	fmt.Println(carFleet(12, position, speed))
}

type pair struct {
	pos int
	sp  int
}

func carFleet(target int, position []int, speed []int) int {
	cnt := 0
	lens := len(position)
	pairs := make([]pair, lens)
	for i := 0; i < len(position); i++ {
		pairs[i] = pair{
			pos: position[i],
			sp:  speed[i],
		}
	}

	sort.SliceStable(pairs, func(i, j int) bool { return pairs[i].pos > pairs[j].pos })

	timeUsed := -1.0
	for i := 0; i < lens; i++ {
		newTimeUsed := float64(target-pairs[i].pos) / float64(pairs[i].sp)
		if (newTimeUsed - timeUsed) > 0.00001 {
			timeUsed = newTimeUsed
			cnt++
		}
	}
	return cnt
}

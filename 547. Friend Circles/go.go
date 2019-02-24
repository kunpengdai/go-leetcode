package main

import "fmt"

func main() {
	fmt.Println(findCircleNum([][]int{{1,1,0}, {1,1,0}, {0,0,1}}))
}
func findCircleNum(M [][]int) int {
	length := len(M)
	ps := make([]int, length)
	for i := 0; i < length; i++ {
		ps[i] = -1
	}
	for i := 0; i < length; i++ {
		if ps[i] != -1 {
			continue
		}
		ps[i] = i
		findNeighber(M, ps, i, i)
	}
	return res(ps)
}
func findNeighber(M [][]int, ps []int, i, num int) {
	for j := 0; j < len(ps); j++ {
		if M[i][j] == 1 && ps[j] == -1 {
			ps[j] = num
			findNeighber(M, ps, j, num)
		}
	}
}

func res(ps []int) int {
	m := make(map[int]bool)
	for _, num := range ps {
		m[num] = true
	}
	return len(m)
}

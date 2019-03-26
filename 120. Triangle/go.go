package Triangle

func main() {

}

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	res := make([]int, len(triangle))

	for i := 0; i < len(triangle); i++ {
		tmp := make([]int, len(triangle))
		for j := 0; j <= i; j++ {
			if j == 0 {
				tmp[0] = res[0] + triangle[i][j]
			} else if j == i {
				tmp[j] = res[j-1] + triangle[i][j]
			} else {
				tmp[j] = min(res[j-1], res[j]) + triangle[i][j]
			}
		}
		res = tmp
	}

	minValue := 0xffffffff
	for _, num := range res {
		if num < minValue {
			minValue = num
		}
	}
	return minValue
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

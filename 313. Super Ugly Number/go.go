package main

import (
	"fmt"
)

func main() {
	fmt.Println(nthSuperUglyNumber(12, []int{2, 7, 13, 19}))
}

func nthSuperUglyNumber(n int, primes []int) int {
	res := []int{1}

	pt := make([]int, len(primes))
	for {
		minV := 0xffffffff
		index := 0
		for i := 0; i < len(primes); i++ {
			if res[pt[i]]*primes[i] < minV {
				minV = res[pt[i]] * primes[i]
				index = i
			}
		}
		if res[len(res)-1] != minV {
			res = append(res, minV)
		}

		pt[index]++
		if len(res) > n {
			break
		}
	}

	return res[n-1]
}

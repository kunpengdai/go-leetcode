package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

func maxProfit(prices []int) int {
	minPrice, maxFit := 0xffffffff, 0
	for _, price := range prices {
		if price-minPrice > maxFit {
			maxFit = price - minPrice
		}
		if price < minPrice {
			minPrice = price
		}
	}
	return maxFit
}

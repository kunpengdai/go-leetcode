package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(multiply("22", "55"))
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	resInt := make([]int, len(num1)+len(num2))
	for i := 0; i < len(num1); i++ {
		for j := 0; j < len(num2); j++ {
			resInt[i+j+1] += int(num1[i]-'0') * int(num2[j]-'0')
		}
	}
	for i := len(resInt) - 1; i > 0; i-- {
		if resInt[i] >= 10 {
			resInt[i-1] += resInt[i] / 10
			resInt[i] = resInt[i] % 10
		}
	}
	if resInt[0] == 0 {
		resInt = resInt[1:]
	}
	res := ""
	for _, v := range resInt {
		res = res + strconv.Itoa(v)
	}
	return res
}

package main

import (
	"math"
)

//func main() {
//	fmt.Println(printNumbers(2))
//}

func printNumbers(n int) []int {
	max := 0
	for i := 0; i < n; i++ {
		max += 9 * int(math.Pow10(i))
	}
	res := make([]int, 0)
	for i := 1; i <= max; i++ {
		res = append(res, i)
	}
	return res
}

package main

import "fmt"

/*
0 - 100
1 2 3 4 5 6 7 8 9 10 11 12 13
*/

func main() {
	fmt.Println(solveZhi(100))
}

func solveZhi(n int) []int {
	isZhi := func(a int) bool {
		for i := 2; i < a; i++ {
			if a%i == 0 {
				return false
			}
		}
		return true
	}

	res := make([]int, 0)
	for i := 2; i <= n; i++ {
		if isZhi(i) {
			res = append(res, i)
		}
	}
	return res
}

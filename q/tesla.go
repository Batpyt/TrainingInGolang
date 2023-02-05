package main

import "fmt"

func main() {
	//test := 10011
	//Solution(test)
	//fmt.Println(reverse(test))

	fmt.Println(subarraySum([]int{2, -2, 3, 0, 4, -7}))
	//fmt.Println(citiesNetworkRank([]int{1, 2, 3, 3}, []int{2, 3, 1, 4}, 4))
}

func Solution(N int) {
	var enable_print int
	enable_print = N % 10
	for N > 0 {
		//fmt.Println(N, N %10, enable_print)
		if enable_print == 0 && N%10 != 0 {
			enable_print = 1
			//fmt.Println("change enable")
		}
		if enable_print == 1 {
			fmt.Print(N%10, "q ")
		}
		N = N / 10
	}
}

func reverse(x int) int {
	var res int
	for x != 0 {
		digit := x % 10
		fmt.Println(x, digit)
		x /= 10
		res = res*10 + digit
	}
	return res
}

func subarraySum(A []int) int {
	var res, preSum int
	exist := make(map[int]int)
	exist[0] = 1
	for _, a := range A {
		preSum += a
		if _, ok := exist[preSum]; ok {
			res += exist[preSum]
		}
		exist[preSum] += 1
		fmt.Println(exist, preSum, res, a)
	}
	return res
}

type city struct {
	neighbors []int
}

func citiesNetworkRank(A []int, B []int, N int) int {
	cities := make([]*city, N+1)
	for i := 1; i <= N; i++ {
		cities[i] = &city{}
	}
	for i := 0; i < len(A); i++ {
		a := A[i]
		b := B[i]
		cities[a].neighbors = append(cities[a].neighbors, b)
		cities[b].neighbors = append(cities[b].neighbors, a)
	}
	res := 0
	for _, c := range cities {
		if c == nil || c.neighbors == nil || len(c.neighbors) == 0 {
			continue
		}
		for _, n := range c.neighbors {
			sum := len(c.neighbors) + len(cities[n].neighbors)
			//fmt.Println(index, n, sum, c, cities[n], len(c.neighbors), len(cities[n].neighbors))
			sum--
			if sum > res {
				res = sum
			}
		}
	}
	return res
}

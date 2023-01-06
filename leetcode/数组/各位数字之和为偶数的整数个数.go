package main

func countEven(num int) int {
	var res int
	for i := 1; i <= num; i++ {
		sum := 0
		temp := i
		for temp > 0 {
			sum += temp % 10
			temp /= 10
		}
		if sum%2 == 0 {
			res++
		}
	}
	return res
}

package main

func main() {

}

func arithmeticTriplets(nums []int, diff int) int {
	values := make(map[int]bool)
	for _, n := range nums {
		values[n] = true
	}

	res := 0
	for _, n := range nums {
		_, ok1 := values[n+diff]
		_, ok2 := values[n+2*diff]
		if ok1 && ok2 {
			res++
		}
	}
	return res
}

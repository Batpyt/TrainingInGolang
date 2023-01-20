package main

import (
	"fmt"
	"strconv"
)

//func main() {
//	fmt.Println(digitCount("1210"))
//	fmt.Println(digitCount("030"))
//}
func digitCount(num string) bool {
	count := make(map[int]int)
	for _, n := range num {
		i, _ := strconv.ParseInt(string(n), 10, 64)
		if _, ok := count[int(i)]; !ok {
			count[int(i)] = 1
			continue
		}
		count[int(i)] = count[int(i)] + 1
	}

	fmt.Println(count)

	for index, n := range num {
		c, _ := strconv.ParseInt(string(n), 10, 64)
		if int(c) != count[index] {
			return false
		}
	}
	return true
}

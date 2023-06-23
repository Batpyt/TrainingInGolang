package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(4)
	fmt.Println(runtime.NumCPU())
	sliceTest()
	arrTest()
}

func sliceTest() {
	s := make([]int, 3, 5)
	fmt.Println("Len:", len(s), "Cap:", cap(s)) // Len: 3 Cap: 5

	s = append(s, 1, 2, 3)
	fmt.Println("Len:", len(s), "Cap:", cap(s)) // Len: 6 Cap: 10

	s = append(s, 4, 5, 6, 7, 8)
	fmt.Println("Len:", len(s), "Cap:", cap(s)) // Len: 11 Cap: 20
}

func arrTest() {
	s := [4]int{1, 2, 3, 4}
	fmt.Println("Len:", len(s), "Cap:", cap(s))

	//s = append(s, 1,2,3,4) //报错，数组不能append
	fmt.Println("Len:", len(s), "Cap:", cap(s))
}

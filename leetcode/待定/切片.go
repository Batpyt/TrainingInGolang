package main

import "fmt"

func main() {
	var s1 []int         //nil 切片
	s2 := []int{}        //空切片
	s3 := make([]int, 0) //空切片

	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)
	fmt.Println(s3 == nil)
}

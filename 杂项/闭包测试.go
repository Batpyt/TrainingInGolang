package main

import "fmt"

func main() {

	i := incr()      //i成了一个闭包，i保存着对x的引用，导致x逃逸了
	fmt.Println(i()) // 别忘记括号，不加括号相当于地址
	fmt.Println(i())

	//下面两次调用直接返回了三个闭包，三个不同的x
	fmt.Println(incr()())
	fmt.Println(incr()())

}

/*
结果
1
2
1
1
*/

func incr() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

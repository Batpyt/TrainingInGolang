package main

import (
	"fmt"
	"sync"
)

func main() {
	test2()
	test3()
}

func test1() {
	total, sum := 0, 0
	for i := 1; i <= 10; i++ {
		sum += i
		go func() {
			total += i
		}()
	}
	fmt.Printf("total: %v, sum: %v", total, sum)
}

/*
total 值不确定的原因是因为在并发执行的匿名函数中，它们共享了相同的变量 i 和 total。

在循环结束后，所有的并发函数才开始执行。此时，i 的值已经是 11。
而在并发函数内部，它们引用的是外部作用域的变量 i 和 total。

由于并发函数的执行是异步的，它们在访问和修改 total 变量时可能会发生竞争条件（race condition）。
多个并发函数同时修改 total 变量可能导致不确定的结果。

具体来说，在并发函数中，每个函数都会读取当前循环结束时的 i 值，并将其添加到 total 中。
由于所有并发函数都是在循环结束后才执行的，它们读取的 i 值都是相同的，即 11。因此，它们都将 11 添加到 total 中。

由于并发函数的执行顺序不确定，不同的并发函数可能以不同的顺序执行，导致 total 的最终值不确定。
每次运行程序时，由于并发执行的不确定性，最终的 total 值可能不同。

要解决这个问题，可以通过在并发函数内部使用参数传递来创建一个局部变量，并将当前循环的值传递给它，
以避免共享变量导致的竞争条件。例如，可以将 i 的值作为参数传递给并发函数，并在函数内部使用该参数进行计算，
从而避免竞争条件和不确定的结果。
*/

func test2() {
	total, sum := 0, 0
	var mutex sync.Mutex // 创建互斥锁

	wg := sync.WaitGroup{}
	wg.Add(10)

	for i := 1; i <= 10; i++ {
		sum += i
		go func(num int) {
			defer wg.Done()
			mutex.Lock() // 加锁
			total += num
			mutex.Unlock() // 解锁
		}(i)
	}

	wg.Wait() // 等待所有 goroutine 完成

	fmt.Printf("total: %v, sum: %v", total, sum)
}

func test3() {
	total, sum := 0, 0
	for i := 1; i <= 10; i++ {
		sum += i
		go func(a int) {
			total += a
		}(i)
	}
	fmt.Printf("total: %v, sum: %v", total, sum)
}

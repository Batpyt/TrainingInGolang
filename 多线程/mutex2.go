package main

import (
	"fmt"
	"sync"
)

// 协程计数器
var wg sync.WaitGroup
var lock sync.Mutex //定义一个互斥锁

var sum int

func main() {
	fmt.Println("嘻嘻")

	wg.Add(2)

	go func() {
		defer wg.Done()

		for i := 0; i < 100000; i++ {
			lock.Lock() // 加锁
			sum += 1
			lock.Unlock() // 解锁
		}
	}()

	go func() {
		defer wg.Done()

		for i := 0; i < 100000; i++ {
			lock.Lock() // 加锁
			sum -= 1
			lock.Unlock() // 解锁
		}
	}()

	wg.Wait() // 等待所有的协程执行完毕

	fmt.Println(sum)
	fmt.Println("主线程退出")
}

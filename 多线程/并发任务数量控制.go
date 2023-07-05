package main

import (
	"fmt"
	"sync"
)

func main() {
	count := 10            // 最大支持并发
	sum := 100             // 任务总数
	wg := sync.WaitGroup{} //控制主协程等待所有子协程执行完之后再退出。

	c := make(chan struct{}, count) // 控制任务并发的chan
	defer close(c)

	for i := 0; i < sum; i++ {
		wg.Add(1)
		c <- struct{}{} // 作用类似于waitgroup.Add(1)
		go func(j int) {
			defer wg.Done()
			fmt.Println(j)
			<-c // 执行完毕，释放资源
		}(i)
	}
	wg.Wait()
}

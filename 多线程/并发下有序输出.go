package main

import (
	"fmt"
	"sync"
)

func main() {
	a := 0
	fn := func() int {
		a++
		return a
	}

	ch := make(chan int, 1)

	var wg sync.WaitGroup

	wg.Add(3)
	for i := 0; i < 3; i++ {
		index := i
		go func() {
			defer wg.Done()
			fmt.Println(index, " begin")
			for {
				ch <- 1
				n := fn()
				if n > 10 {
					break
				}
				fmt.Println(index, n)
				<-ch
			}

			fmt.Println(index, " done")
		}()
	}

	wg.Wait()
	close(ch)
}

// 使用管道来塞其他协程来实现抢占，使用休眠实现安全退出，这种方案是可靠的，同理，最好能够使用WaitGroup来实现安全退出

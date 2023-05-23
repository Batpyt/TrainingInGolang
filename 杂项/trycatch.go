package main

import (
	"fmt"
	"log"
	"time"
)

func chanAndRecover() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
	}()

	for i := 0; i < 5; i++ {
		go func(i int) {
			//如果不加recover来处理panic，一旦一个goroutine发生错误，会影响其他goroutine执行
			defer func() {
				if err := recover(); err != nil { //catch操作
					log.Println("Error:", i)
				}
			}()
			fmt.Println(i, 10/i)
		}(<-ch)
	}

	time.Sleep(2 * time.Second)
}

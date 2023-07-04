package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

func main() {
	chanShop := make(chan string, 10)
	chanTel := make(chan int)

	go Producer(chanShop, chanTel)
	go Consumer(chanShop, chanTel)
	for {
		time.Sleep(time.Second)
	}
}

func Producer(chanShop chan string, chanTel chan int) {
	runtime.GOMAXPROCS(2)
	for i := 0; i < 10; i++ {
		product := strconv.Itoa(time.Now().Nanosecond())
		chanShop <- "产品" + product
		fmt.Println("生产了产品", product)
		time.Sleep(time.Second)
	}
	close(chanShop)
	fmt.Println("生产完毕")
	chanTel <- 123456789
	fmt.Println(123456789, "呼出电话")
}

func Consumer(chanShop chan string, chanTel chan int) {
	runtime.GOMAXPROCS(1)

	tel := <-chanTel
	fmt.Println("收到来电", tel)
	for product := range chanShop {
		fmt.Println("消费了产品", product)
	}
	fmt.Println("消费完毕")
}

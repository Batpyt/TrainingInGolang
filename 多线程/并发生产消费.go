package main

import (
	"fmt"
	"strconv"
)

/*
用两个channel实现并发生产消费模式
*/
var (
	//传递数据
	chanStore = make(chan string, 100)

	//用来判断是否消费完毕
	done = make(chan bool, 100)
)

func main() {
	//生产数据数量
	maxPro := 100
	//消费者数量
	maxCon := 3

	//并发消费者
	for i := 0; i < maxCon; i++ {
		go consumer2(i)
	}

	go producer2(maxPro)

	//判断消费者是否都消费完毕，只有当done传入了maxCon次数据的时候这个循环才会结束
	for i := 1; i <= maxCon; i++ {
		<-done
	}
}

func producer2(max int) {

	//生产数据
	for i := 0; i < max; i++ {
		product := strconv.FormatInt(int64(i), 10)
		chanStore <- "product: " + product

	}
	close(chanStore)
}

func consumer2(index int) {
	//消费数据
	for true {
		product, ok := <-chanStore
		if ok {
			fmt.Println(index, " 消费了 ", product)
		} else {
			break
		}
	}

	//当这个携程拿不到数据了，视作消费结束，往done里传入一个数据
	done <- true
	fmt.Println(index, "is finished!!!!")
}

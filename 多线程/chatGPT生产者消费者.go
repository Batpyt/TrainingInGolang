package main

import (
	"fmt"
	"math/rand"

	"sync"
	"time"
)

const (
	NumProducers = 3
	NumConsumers = 2
	BufferSize   = 5
	NumItems     = 10
)

var wgg sync.WaitGroup

//func main() {
//	buffer := make(chan int, BufferSize)
//	wg.Add(NumProducers + NumConsumers)
//
//	// 启动生产者
//	for i := 1; i <= NumProducers; i++ {
//		go producer(i, buffer)
//	}
//
//	// 启动消费者
//	for i := 1; i <= NumConsumers; i++ {
//		go consumer(i, buffer)
//	}
//
//	wg.Wait()
//	fmt.Println("所有生产者和消费者任务已完成")
//}

func producer1(id int, buffer chan<- int) {
	defer wgg.Done()

	for i := 1; i <= NumItems; i++ {
		// 模拟生产耗时
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

		item := rand.Intn(100)
		buffer <- item
		fmt.Printf("生产者 %d 生产了 %d\n", id, item)
	}

	fmt.Printf("生产者 %d 完成生产\n", id)
}

func consumer1(id int, buffer <-chan int) {
	defer wgg.Done()

	for i := 1; i <= NumItems; i++ {
		item := <-buffer
		fmt.Printf("消费者 %d 消费了 %d\n", id, item)

		// 模拟消费耗时
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}

	fmt.Printf("消费者 %d 完成消费\n", id)
}

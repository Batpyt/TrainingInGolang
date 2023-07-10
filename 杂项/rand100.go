package main

import (
	"fmt"
	"sync"
)

//n 个协程打印n个随机数
//func main() {
//	n := 10
//	ch := make(chan int)
//	arr := make([]int, n)
//	for i := 0; i < n; i++ {
//		arr[i] = rand.Intn(10)
//	}
//
//	wgPro := &sync.WaitGroup{}
//	wgCons := &sync.WaitGroup{}
//
//	wgPro.Add(1)
//	go producer(arr, ch, wgPro)
//
//	for i := 0; i < n; i++ {
//		wgCons.Add(1)
//		go consumer(i, ch, wgCons)
//	}
//
//	wgPro.Wait()
//	close(ch)
//	wgCons.Wait()
//}

func producer(arr []int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < len(arr); i++ {
		ch <- arr[i]
	}
}

func consumer(goroutine int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range ch {
		fmt.Println(num, " no: ", goroutine, " goroutine")
	}
}

package main

import (
	"fmt"
	"sync"
)

var arr = make([]int, 0)

func main() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch3 := make(chan int, 1)
	exit := make(chan int, 1)

	var wg sync.WaitGroup
	wg.Add(3)

	ch1 <- 1
	go func1(ch1, ch2, exit, &wg)
	go func2(ch2, ch3, exit, &wg)
	go func3(ch3, exit, &wg)

	wg.Wait()
	close(ch1)
	close(ch2)
	close(ch3)
	close(exit)
	fmt.Println(arr)
}

func func1(func1Chan, func2Chan, exit chan int, wg *sync.WaitGroup) {
	fmt.Println("func1 running")
	defer wg.Done()

	select {
	case <-func1Chan:
		arr = append(arr, 1, 2, 3, 4)
		func2Chan <- 1
	case <-exit:
		return
	}

}

func func2(func2Chan, func3Chan, exit chan int, wg *sync.WaitGroup) {

	fmt.Println("func2 running")
	defer wg.Done()

	select {
	case <-func2Chan:
		arr = append(arr, 5, 6, 7, 8)
		func3Chan <- 1
	case <-exit:
		return
	}

}

func func3(func3Chan, exit chan int, wg *sync.WaitGroup) {

	fmt.Println("func3 running")
	defer wg.Done()

	select {
	case <-func3Chan:
		arr = append(arr, 9, 10)
		exit <- 1
	case <-exit:
		return
	}

}

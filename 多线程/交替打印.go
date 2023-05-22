package main

import (
	"fmt"
	"sync"
)

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []string{"a", "b", "c", "d", "e"}

	//1a2b3c4d5e

	chan1 := make(chan int)
	chan2 := make(chan int)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for _, a := range arr1 {
			fmt.Print(a)
			chan1 <- 0
			<-chan2
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, a := range arr2 {
			<-chan1
			fmt.Print(a)
			chan2 <- 0
		}
	}()

	wg.Wait()
}

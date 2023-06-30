package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	len := 10
	var wg sync.WaitGroup
	wg.Add(len)
	ch := make(chan int, 10)

	for i := 0; i < len; i++ {
		go insert(ch, &wg)
	}

	wg.Wait()
	close(ch)

	for n := range ch {
		fmt.Println(n)
	}
}

func insert(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- rand.Int()
}

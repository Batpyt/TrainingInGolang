package main

import (
	"math/rand"
	"sync"
)

func main() {
	ch := make(chan int, 10)

	var wg sync.WaitGroup
	wg.Add(1)
	randInt(ch, wg)
	wg.Wait()

	for i := <- ch; i != 0
}

func randInt(ch chan int, wg sync.WaitGroup) {
	defer wg.Done()
	random := rand.Int()
	ch <- random
}
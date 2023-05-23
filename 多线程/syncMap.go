package main

import (
	"fmt"
	"sync"
)

func process(wg *sync.WaitGroup) {
	fmt.Println("start goroutine")
	fmt.Printf("goroutine end\n")
	wg.Done()
}

//func main() {
//	var wg sync.WaitGroup
//	for i := 0; i < 3; i++ {
//		wg.Add(1)
//		go process(&wg)
//	}
//	wg.Wait()
//	fmt.Println("all finished")
//}

func syncMap() {
	var sMap sync.Map
	sMap.Load("null")
}

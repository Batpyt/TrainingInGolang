package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	ch2 := make(chan int, 1)
	go func() {
		ch <- 1
		ch2 <- 2

		close(ch)
		close(ch2)
	}()

	time.Sleep(2 * time.Second)

	v, ok := <-ch
	fmt.Println(v, ok)
	v, ok = <-ch
	fmt.Println(v, ok)

	v2, ok := <-ch2
	fmt.Println(v2, ok)
	v2, ok = <-ch2
	fmt.Println(v2, ok)
}

package main

import (
	"fmt"
	"sync"
)

/*
1. users: id, name; orders: id, user_id, price
*/

/*
select users.name users inner join (select user_id, count(price) as c from orders group by user_id) as table on user_id order by table.c
*/

/*

 */

//func main() {
//	wgProducer := &sync.WaitGroup{}
//	wgConsumer := &sync.WaitGroup{}
//
//	chCount := 10
//	ch := make(chan int, chCount)
//	wgProducer.Add(1)
//	go producer(chCount, ch, wgProducer)
//
//	for i := 0; i < 10; i++ {
//		wgConsumer.Add(1)
//		go consumer(i, ch, wgConsumer)
//	}
//
//	wgProducer.Wait()
//	close(ch)
//	wgConsumer.Wait()
//}

func producer(count int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < count; i++ {
		ch <- i
	}
}

func consumer(index int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range ch {
		fmt.Println(num, " from ", index)
	}
}

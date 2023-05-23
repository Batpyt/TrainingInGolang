package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	ch := make(chan string)
	ch2 := make(chan int)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		who := "外卖小哥我："
		food := "鸡腿"
		fmt.Println(who, "送餐中……2s")

		if random := rand.Intn(10); random < 5 {
			fmt.Println(who, "已送餐到门口，等待顾客开门取餐")
			ch <- food
			fmt.Println(who, "订单已送达，开始送其他单")
		} else {
			fmt.Println(who, "撞车了")
			ch2 <- 1
		}

	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		who := "你："
		fmt.Println(who, "等待外卖……")
		select {
		case food := <-ch:
			fmt.Println(who, "磨磨唧唧开门中……3s")
			fmt.Println(who, "拿到", food, "开吃！")
		case <-ch2:
			fmt.Println(who, "饿死了，差评")
		}

	}()

	wg.Wait()

	close(ch)
	close(ch2)
}

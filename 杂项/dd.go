package main

import (
	"math/rand"
)

/*
发红包算法：总计发m元，给n个人发，每个人随机金额；写一个算法，每个人获得的红包金额。
*/

//func main() {
//	pack := RedPack{
//		Money:      1000,
//		PackCap:    10,
//		MoneySpent: 0,
//		PackSpent:  0,
//	}
//
//	users := make([]int, pack.PackCap)
//	for i := 0; i < pack.PackCap; i++ {
//		users[i] = pack.RandomRedPack()
//		fmt.Println(users, pack)
//	}
//	fmt.Println(users)
//}

type RedPack struct {
	Money   int //分
	PackCap int

	MoneySpent int
	PackSpent  int
}

func (p *RedPack) RandomRedPack() int {
	if p.PackSpent >= p.PackCap {
		return 0
	}

	if p.PackSpent == p.PackCap-1 {
		p.PackSpent++
		money := p.Money - p.MoneySpent
		p.MoneySpent = p.Money
		return money
	}

	haveToLeftMoney := p.PackCap - p.PackSpent

	amount := rand.Intn(p.Money - p.MoneySpent - haveToLeftMoney)
	if amount <= 0 {
		amount++
	}
	p.MoneySpent += amount
	p.PackSpent++
	return amount
}

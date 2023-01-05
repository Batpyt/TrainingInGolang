package main

import "sync"

type Bank struct {
	lock    sync.Mutex
	account int
}

func (b *Bank) In(money int) {
	b.lock.Lock()
	defer b.lock.Unlock()
	b.account += money
}

func (b *Bank) Out(money int) {
	b.lock.Lock()
	defer b.lock.Unlock()
	b.account -= money
}

package main

import (
	"fmt"
	"sync"
)

type Counter interface {
	Inc() uint64
	Dec() uint64
	Cur() uint64
}

type MyCounter struct {
	lock  sync.Mutex
	count uint64
}

func (m *MyCounter) Inc() uint64 {
	m.lock.Lock()
	m.count++
	defer m.lock.Unlock()
	fmt.Println(m.count)
	return m.count
}

func (m *MyCounter) Dec() uint64 {
	m.lock.Lock()
	m.count--
	defer m.lock.Unlock()
	if m.count <= 0 {
		m.count = 0
		return 0
	}
	return m.count
}

func (m *MyCounter) Cur() uint64 {
	m.lock.Lock()
	res := m.count
	defer m.lock.Unlock()
	return res
}

func main() {
	var c Counter = &MyCounter{}

	c.Dec()
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc()
		}()
	}

	wg.Wait()

	fmt.Println(c.Cur())
}

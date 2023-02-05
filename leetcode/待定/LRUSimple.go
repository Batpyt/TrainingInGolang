package main

import (
	"container/list"
	"fmt"
)

func main() {
	c := Construct(2)
	fmt.Println(c)
	c.put(1, 1)
	fmt.Println(c)
	c.put(2, 2)
	fmt.Println(c)
	fmt.Println(c.get(1))
	fmt.Println(c)
	c.put(3, 3)
	fmt.Println(c)
}

type LRUCacheSimple struct {
	size, capacity int
	cache          map[interface{}]*list.Element
	l              *list.List
}

type node struct {
	key, value interface{}
}

func Construct(cap int) *LRUCacheSimple {
	return &LRUCacheSimple{
		size:     0,
		capacity: cap,
		cache:    map[interface{}]*list.Element{},
		l:        list.New(),
	}
}

func (l *LRUCacheSimple) get(key interface{}) (interface{}, bool) {
	if _, ok := l.cache[key]; !ok {
		fmt.Println("nil")
		return nil, false
	}
	l.l.MoveToFront(l.cache[key])
	return l.cache[key].Value.(*node).value, true
}

func (l *LRUCacheSimple) put(key, value interface{}) {
	if _, ok := l.cache[key]; ok {
		l.cache[key].Value.(*node).value = value
		l.l.MoveToFront(l.cache[key])
	} else {
		e := l.l.PushFront(&node{
			key:   key,
			value: value,
		})
		l.cache[key] = e
		if l.size+1 > l.capacity {
			if tail := l.l.Back(); tail != nil {
				delete(l.cache, tail.Value.(*node).key)
				l.l.Remove(tail)
			}
		}
		l.size++
	}
}

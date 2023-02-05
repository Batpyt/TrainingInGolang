package main

//func main() {
//	c := Constructor(2)
//	fmt.Println(c)
//	c.Put(1, 1)
//	fmt.Println(c.cache)
//	c.Put(2, 2)
//	fmt.Println(c.cache)
//	c.Get(1)
//	fmt.Println(c.cache)
//	c.Put(3, 3)
//	fmt.Println(c.cache)
//}

/*
通过双向链表实现key的排序，kv值存在链表的节点中
*/

type LRUCache struct {
	size     int                  //已存的数量
	capacity int                  //最大容量
	cache    map[int]*DLinkedNode //存放key和链表节点
	//链表的头尾都为dummy node
	head *DLinkedNode //链表的头
	tail *DLinkedNode //链表尾
}

/*
链表中存放kv和前后节点
存key的用处是超过容量删除队尾元素时通过key来删除map中元素
*/

type DLinkedNode struct {
	key   int
	value int
	pre   *DLinkedNode
	next  *DLinkedNode
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		size:     0,
		capacity: capacity,
		cache:    map[int]*DLinkedNode{},
		head: &DLinkedNode{
			key:   0,
			value: 0,
			pre:   nil,
			next:  nil,
		},
		tail: &DLinkedNode{
			key:   0,
			value: 0,
			pre:   nil,
			next:  nil,
		},
	}

	//初始状态，首尾相连
	l.head.next = l.tail
	l.tail.pre = l.head
	return l
}

/*
没有查询到返回-1，查询到了将其移到队首
*/

func (this *LRUCache) Get(key int) int {
	v, ok := this.cache[key]
	if !ok {
		return -1
	}
	this.moveToHead(v)
	return v.value
}

func (this *LRUCache) Put(key int, value int) {
	//若已存在，直接替换value值
	if _, ok := this.cache[key]; ok {
		this.cache[key].value = value
		this.moveToHead(this.cache[key])
	} else {
		//是新值时，添加到队首
		node := &DLinkedNode{
			key:   key,
			value: value,
			pre:   nil,
			next:  nil,
		}
		this.addToHead(node)
		//并判断是否超过容量上线，移除队尾元素
		if this.size+1 > this.capacity {
			removeKey := this.tail.pre.key
			this.removeNode(this.tail.pre)
			delete(this.cache, removeKey)
		}
		this.size++
		this.cache[key] = node
	}
}

func (this *LRUCache) removeNode(n *DLinkedNode) {
	n.pre.next = n.next
	n.next.pre = n.pre
}

func (this *LRUCache) addToHead(n *DLinkedNode) {
	n.pre = this.head
	n.next = this.head.next
	this.head.next.pre = n
	this.head.next = n
}

func (this *LRUCache) moveToHead(n *DLinkedNode) {
	this.removeNode(n)
	this.addToHead(n)
}

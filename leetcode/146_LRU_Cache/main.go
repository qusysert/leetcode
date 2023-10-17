// https://leetcode.com/problems/lru-cache/
package main

func main() {
	lru := Constructor(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	lru.Get(1)
	lru.Put(3, 3)
	lru.Get(2)
	lru.Put(4, 4)
	lru.Get(1)
	lru.Get(3)
	lru.Get(4)
}

type Node struct {
	next, prev *Node
	key        int
	value      int
}

type LRUCache struct {
	capacity   int
	cache      map[int]*Node
	head, tail *Node
}

func Constructor(capacity int) LRUCache {
	head := Node{key: 0, value: 0, prev: nil, next: nil}
	tail := Node{key: 0, value: 0, prev: nil, next: nil}
	head.next = &tail
	tail.prev = &head
	return LRUCache{capacity: capacity, cache: make(map[int]*Node), head: &head, tail: &tail}
}

func (this *LRUCache) Insert(node *Node) {
	// left side of node
	node.prev = this.tail.prev
	this.tail.prev.next = node

	// right side of node
	node.next = this.tail
	this.tail.prev = node
}

func (this *LRUCache) Remove(node *Node) {
	prev := node.prev
	next := node.next

	prev.next = node.next
	next.prev = node.prev
}

func (this *LRUCache) Get(key int) int {
	val, ok := this.cache[key]
	if ok {
		this.Remove(val)
		this.Insert(val)
		return val.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	node := &Node{key: key, value: value}

	// if this key already exists - update pointer in the cache map
	if _, ok := this.cache[key]; ok {
		this.Remove(this.cache[key])
		this.Insert(node)
		this.cache[key] = node
		return
	}

	// if cache is overflown
	if len(this.cache)+1 > this.capacity {
		delete(this.cache, this.head.next.key)
		this.Remove(this.head.next)
		this.Insert(node)
		this.cache[node.key] = node
		return
	}

	// if new key and we have capacity
	this.cache[key] = node
	this.Insert(node)
}

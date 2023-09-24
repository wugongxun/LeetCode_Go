package main

func main() {

}

type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*LinkedNode
	head, tail *LinkedNode
}

//func Constructor(capacity int) LRUCache {
//	lru := LRUCache{
//		capacity: capacity,
//		cache:    map[int]*LinkedNode{},
//		head:     initLinkedNode(0, 0),
//		tail:     initLinkedNode(0, 0),
//	}
//	lru.head.next = lru.tail
//	lru.tail.prev = lru.head
//	return lru
//}

func (lru *LRUCache) Get(key int) int {
	if _, ok := lru.cache[key]; !ok {
		return -1
	}
	node := lru.cache[key]
	lru.moveToHead(node)
	return node.value
}

func (lru *LRUCache) Put(key int, value int) {
	if _, ok := lru.cache[key]; !ok {
		node := initLinkedNode(key, value)
		lru.cache[key] = node
		lru.addToHead(node)
		lru.size++
		if lru.size > lru.capacity {
			removed := lru.removeTail()
			delete(lru.cache, removed.key)
			lru.size--
		}
	} else {
		node := lru.cache[key]
		node.value = value
		lru.moveToHead(node)
	}
}

func (lru *LRUCache) addToHead(node *LinkedNode) {
	node.prev = lru.head
	node.next = lru.head.next
	lru.head.next.prev = node
	lru.head.next = node
}

func (lru *LRUCache) removeNode(node *LinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (lru *LRUCache) moveToHead(node *LinkedNode) {
	lru.removeNode(node)
	lru.addToHead(node)
}

func (lru *LRUCache) removeTail() *LinkedNode {
	node := lru.tail.prev
	lru.removeNode(node)
	return node
}

type LinkedNode struct {
	key, value int
	prev, next *LinkedNode
}

func initLinkedNode(key, value int) *LinkedNode {
	return &LinkedNode{key: key, value: value}
}

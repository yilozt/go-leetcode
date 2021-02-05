/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU 缓存机制
 */

// @lc code=start
package main

// 链表节点
type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

// 双向链表
type DoubleList struct {
	begin *Node
	end   *Node
}

func (d *DoubleList) init() {
	d.begin = &Node{}
	d.end = &Node{}
	d.begin.next = d.end
	d.end.prev = d.begin
}

func (d *DoubleList) insert(key, value int) *Node {
	tmp := d.begin.next
	nodeToInsert := &Node{
		key:   key,
		value: value,
		prev:  d.begin,
		next:  tmp,
	}
	tmp.prev = nodeToInsert
	d.begin.next = nodeToInsert
	return nodeToInsert
}

func (d *DoubleList) delete(node *Node) {
	prev := node.prev
	next := node.next
	prev.next, next.prev = next, prev
}

func (d *DoubleList) getLast() *Node {
	return d.end.prev
}

func (d *DoubleList) moveToFirst(node *Node) {
	// 首先删掉这个节点
	d.delete(node)
	//接着把节点放到最前面
	tmp := d.begin.next
	node.prev, node.next = d.begin, tmp
	d.begin.next, tmp.prev = node, node
}

type LRUCache struct {
	length   int
	capacity int
	cache    map[int]*Node
	d        DoubleList
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		cache:    make(map[int]*Node),
		d:        DoubleList{},
		capacity: capacity,
	}
	lru.d.init()
	return lru
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.d.moveToFirst(node)
		return node.value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		node.value = value
		this.d.moveToFirst(node)
	} else {
		if this.length < this.capacity {
			this.length++
		} else {
			last := this.d.getLast()
			this.d.delete(last)
			delete(this.cache, last.key)
		}
		newNode := this.d.insert(key, value)
		this.cache[key] = newNode
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

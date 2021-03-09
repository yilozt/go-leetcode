package main

import "container/heap"

type MinHeap []int

type MaxHeap []int

func (m *MinHeap) Push(x interface{}) {
	*m = append(*m, x.(int))
}
func (m *MinHeap) Pop() interface{} {
	last := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return last
}

func (m MinHeap) Len() int {
	return len(m)
}

func (m MinHeap) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m MinHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MaxHeap) Push(x interface{}) {
	*m = append(*m, x.(int))
}
func (m *MaxHeap) Pop() interface{} {
	last := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return last
}

func (m MaxHeap) Len() int {
	return len(m)
}

func (m MaxHeap) Less(i, j int) bool {
	return m[i] > m[j]
}

func (m MaxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

type MedianFinder struct {
	max MaxHeap
	min MinHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	result := MedianFinder{}
	// 初始化大顶堆和小顶堆
	// 大顶堆中存储数组中小的一半，小顶堆中存储数组中大的一半
	heap.Init(&(result.min))
	heap.Init(&(result.max))
	return result
}

func (this *MedianFinder) AddNum(num int) {
	// 大顶堆
	heap.Push(&this.max, num)
	// 大顶堆中最大的送到小顶堆
	heap.Push(&this.min, heap.Pop(&this.max))
	// 如果两个堆内元素相差大于1
	if len(this.max)+1 < len(this.min) {
		heap.Push(&this.max, heap.Pop(&this.min))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	// 如果总元素数量为奇数
	if len(this.min) > len(this.max) {
		return float64(this.min[0])
	}
	return float64(this.min[0]+this.max[0]) / 2
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

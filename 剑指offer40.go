package main

import (
	"container/heap"
)

type IntHeap []int

func (ih IntHeap) Len() int {
	return len(ih)
}

func (ih IntHeap) Less(i, j int) bool {
	return ih[i] > ih[j]
}

func (ih IntHeap) Swap(i, j int) {
	ih[i], ih[j] = ih[j], ih[i]
}

func (ih *IntHeap) Push(i interface{}) {
	*ih = append(*ih, i.(int))
}

func (ih *IntHeap) Pop() interface{} {
	result := (*ih)[len(*ih)-1]
	*ih = (*ih)[:len(*ih)-1]
	return result
}

func getLeastNumbers(arr []int, k int) []int {
	// 可以使用快排中的基准函数解决该问题，参见剑指offer39-1
	// 该处使用辅助大顶堆解决该问题
	if k < 1 || len(arr) == 0 {
		return nil
	}
	var h IntHeap
	heap.Init(&h)
	for _, item := range arr {
		if h.Len() < k {
			heap.Push(&h, item)
		} else {
			if item < h[0] {
				heap.Pop(&h)
				heap.Push(&h, item)
			}
		}
	}
	return h
}

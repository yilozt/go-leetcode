/*
 * @lc app=leetcode.cn id=1675 lang=golang
 *
 * [1675] 数组的最小偏移量
 */

// @lc code=start
package main

import (
	"container/heap"
	"math"
)

type IntHeap []int

/*
Golang中堆需要实现的接口：
type Interface interface {
    Len() int
    Less(i, j int) bool
    Swap(i, j int)
    Push(x interface{}) // add x as element Len()
    Pop() interface{}   // remove and return element Len() - 1.
}
*/

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minimumDeviation(nums []int) int {
	minValue := math.MaxInt32
	// 初始化优先队列
	priorityQueue := &IntHeap{}
	heap.Init(priorityQueue)
	// 首先将所有奇数乘以二变为偶数，将值放入优先队列并在过程中记录最小值
	for index, value := range nums {
		if isOddNum(value) {
			nums[index] = value * 2
		}
		minValue = min(minValue, nums[index])
		heap.Push(priorityQueue, nums[index])
	}
	// 最终结果
	result := math.MaxInt32
	// 按照优先队列顺序依次弹出
	for priorityQueue.Len() > 0 {
		maxValue := heap.Pop(priorityQueue).(int)
		// 计算当前偏移量
		result = min(result, maxValue-minValue)
		if isOddNum(maxValue) {
			break
		}
		maxValue >>= 1
		minValue = min(minValue, maxValue)
		heap.Push(priorityQueue, maxValue)
	}
	return result
}

// 判断是否是奇数
func isOddNum(num int) bool {
	if num%2 == 0 {
		return false
	}
	return true
}

// 返回两个数较小的一个
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end

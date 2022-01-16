package main

import "fmt"

// 有两个排序的数组A1和A2，内存在A1的末尾有足够多的空余空间容纳A2。
// 请实现一个函数，把A2中的所有数字插入A1中，并且所有的数字是排序的。

func main() {
	A1 := make([]int, 6)
	A1[0] = 0
	A1[1] = 2
	A1[2] = 5
	A2 := []int{1, 3, 4}
	merge(A1, A2, 3, 3)
	fmt.Println(A1)
}

func merge(A1, A2 []int, lenA1, lenA2 int) {
	// 基本思路是归并，但如果从前向后归并，会有重复移动的开销，因此考虑从后向前归并
	startA1, startA2 := lenA1-1, lenA2-1
	startResult := lenA1 + lenA2 - 1
	for startA1 >= 0 && startA2 >= 0 {
		if A1[startA1] > A2[startA2] {
			A1[startResult] = A1[startA1]
			startA1--
		} else {
			A1[startResult] = A2[startA2]
			startA2--
		}
		startResult--
	}
}

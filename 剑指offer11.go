package main

func minArray(numbers []int) int {
	low := 0
	high := len(numbers) - 1
	for low < high {
		pivot := low + (high-low)/2
		if numbers[pivot] < numbers[high] {
			// 中值位于右边的递增序列，向左折半
			// 这里不肯定pivot是否是最小值，所以查找范围内需要包括pivot
			high = pivot
		} else if numbers[pivot] > numbers[high] {
			// 中值位于左边的递增序列，向右折半
			low = pivot + 1
		} else {
			// 无法判断位于哪边的序列，线性查找
			high--
		}
	}
	return numbers[low]
}

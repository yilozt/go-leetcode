package main

func QuickSort(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, start, end int) {
	if start < end {
		pivot := partition(nums, start, end)
		quickSort(nums, start, pivot-1)
		quickSort(nums, pivot+1, end)
	}
}

func partition(nums []int, start, end int) int {
	toCompare := nums[end]
	for i := start; i < end; i++ {
		if nums[i] < toCompare {
			swap(nums, i, start)
			start++
		}
	}
	swap(nums, start, end)
	return start
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

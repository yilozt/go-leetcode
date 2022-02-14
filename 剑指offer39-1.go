package main

func majorityElement(nums []int) int { // 如果某个数字超过数组元素一半，那中位数必定是它，基于此可使用快排思路
	if len(nums) == 0 {
		return 0
	}
	start, end := 0, len(nums)-1
	middle := len(nums) >> 1
	pivot := partition(nums, start, end)
	for pivot != middle {
		if pivot < middle {
			start = pivot + 1
			pivot = partition(nums, start, end)
		} else {
			end = pivot - 1
			pivot = partition(nums, start, end)
		}
	}
	return nums[middle]
}

func partition(nums []int, start, end int) int {
	num := nums[end]
	for i := start; i <= end; i++ {
		if nums[i] < num {
			nums[i], nums[start] = nums[start], nums[i]
			start++
		}
	}
	nums[start], nums[end] = nums[end], nums[start]
	return start
}

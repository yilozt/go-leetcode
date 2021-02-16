package main

import "fmt"

func main() {
	test := []int{1, 3, 5, 8, 10, 2}
	fmt.Println(heapSort(test))
}

func heapSort(nums []int) []int {
	// 构建大顶堆
	for i := len(nums)/2 - 1; i >= 0; i-- {
		sink(nums, i, len(nums))
	}
	for i := len(nums) - 1; i >= 1; i-- {
		swap(nums, 0, i)
		sink(nums, 0, i)
	}
	return nums
}

func sink(nums []int, i, length int) {
	left, right := i*2+1, i*2+2
	flag := i
	if left < length && nums[left] > nums[flag] {
		flag = left
	}
	if right < length && nums[right] > nums[flag] {
		flag = right
	}
	if flag == i {
		return
	}
	swap(nums, flag, i)
	sink(nums, flag, length)

}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

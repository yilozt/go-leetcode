package main

func exchange(nums []int) []int {
	// 个人思路是双指针，遇到奇数就放到前面
	if len(nums) == 0 {
		return nums
	}
	swap := func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	}
	left, right := 0, 0
	for right < len(nums) {
		if nums[right]&1 == 1 {
			swap(left, right)
			left++
		}
		right++
	}
	return nums
}

/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

package main

import (
	"fmt"
	"reflect"
)

// @lc code=start
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	l := binarySearchL(nums, target)
	r := binarySearchR(nums, target)

	return []int{l, r}
}

// 寻找左边界
func binarySearchL(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for {
		mid := (l + r) / 2
		if nums[mid] == target && (mid == l || nums[mid-1] != target) {
			return mid
		}
		if l > r {
			return -1
		}

		if target <= nums[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

}

// 寻找右边界
func binarySearchR(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for {
		mid := (l + r) / 2
		if nums[mid] == target && (mid == r || nums[mid+1] != target) {
			return mid
		}
		if l > r {
			return -1
		}

		if target < nums[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
}

// @lc code=end

func main() {
	datas := []struct {
		input  []int
		ans    []int
		target int
	}{
		{input: []int{5, 7, 7, 8, 8, 10}, ans: []int{3, 4}, target: 8},
		{input: []int{5, 7, 7, 8, 8, 10}, ans: []int{-1, -1}, target: 6},
		{input: []int{}, ans: []int{-1, -1}, target: 0},
		{input: []int{1}, ans: []int{0, 0}, target: 1},
	}

	for _, data := range datas {
		got := searchRange(data.input, data.target)
		want := data.ans

		if !reflect.DeepEqual(got, want) {
			fmt.Printf("got: %v, want: %v", got, want)
		}
	}
}

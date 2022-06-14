/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

package main

// @lc code=start
func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int, 0)

	for i, num := range nums {
		rest, ok := hashMap[target-num]
		if ok {
			return []int{i, rest}
		}
		hashMap[num] = i
	}

	return make([]int, 0)
}

// @lc code=end

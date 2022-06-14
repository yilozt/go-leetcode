/*
 * @lc app=leetcode.cn id=136 lang=golang
 *
 * [136] 只出现一次的数字
 */

package main

// @lc code=start
func singleNumber(nums []int) int {
	hash_map := make(map[int]int, 0)
	for _, n := range nums {
		_, ok := hash_map[n]
		if ok {
			delete(hash_map, n)
		} else {
			hash_map[n] = 1
		}
	}
	for n, _ := range hash_map {
		return n
	}
	return 0
}

// @lc code=end

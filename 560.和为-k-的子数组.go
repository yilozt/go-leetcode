/*
 * @lc app=leetcode.cn id=560 lang=golang
 *
 * [560] 和为 K 的子数组
 */

package main

// @lc code=start
func subarraySum(nums []int, k int) int {
	len := len(nums)
	sums := make([]int, len)

	{
		sum := 0
		for i, n := range nums {
			sum += n
			sums[i] = sum
		}
	}

	count := 0

	for i := 0; i < len; i++ {
		for j := i; j < len; j++ {
			sum := sums[j]
			if i-1 >= 0 {
				sum -= sums[i-1]
			}
			if sum == k {
				count++
			}
		}
	}

	return count
}

// @lc code=end

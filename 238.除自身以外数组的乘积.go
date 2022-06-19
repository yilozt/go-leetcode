/*
 * @lc app=leetcode.cn id=238 lang=golang
 *
 * [238] 除自身以外数组的乘积
 */

package main

// @lc code=start
func productExceptSelf(nums []int) []int {
	len := len(nums)
	ans := make([]int, len)
	for i := 0; i < len; i++ {
		ans[i] = 1
	}
	prefix, suffix := 1, 1

	pre, suf := 0, len-1
	for pre < len {
		if pre >= 1 {
			prefix *= nums[pre-1]
		}
		if suf <= len-2 {
			suffix *= nums[suf+1]
		}

		ans[pre] *= prefix
		ans[suf] *= suffix

		pre++
		suf--
	}

	return ans
}

// @lc code=end

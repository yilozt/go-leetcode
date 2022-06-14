/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */

package main

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	a := make([]int, m)
	b := nums2
	res := nums1
	copy(a, res[0:m])

	i_a, i_b, i := 0, 0, 0
	for i_a < m && i_b < n {
		if a[i_a] < b[i_b] {
			res[i] = a[i_a]
			i_a++
		} else {
			res[i] = b[i_b]
			i_b++
		}
		i++
	}

	for i_a < m {
		res[i] = a[i_a]
		i++
		i_a++
	}

	for i_b < n {
		res[i] = b[i_b]
		i++
		i_b++
	}
}

// @lc code=end

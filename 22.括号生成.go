/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start
package main

func generateParenthesis(n int) []string {
	var res []string
	backtrace("", n, n, &res)
	return res
}

func backtrace(current string, left, right int, res *[]string) {
	if left == 0 && right == 0 {
		*res = append(*res, current)
	}
	if left > 0 {
		backtrace(current+"(", left-1, right, res)
	}
	if left < right {
		backtrace(current+")", left, right-1, res)
	}
}

// @lc code=end

/*
 * @lc app=leetcode.cn id=72 lang=golang
 *
 * [72] 编辑距离
 */

// @lc code=start
package main

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	// m+1行n+1列矩阵
	helper := make([][]int, m+1)
	for index := range helper {
		helper[index] = make([]int, n+1)
	}
	for i := 0; i <= m; i++ {
		// 长度为i的字符串到长度为0的字符串需要删除i次
		helper[i][0] = i
	}
	for j := 1; j <= n; j++ {
		// 长度为0的字符串到长度为j的字符串需要插入j次
		helper[0][j] = j
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				// 如果该位相同，则变动次数等于去掉这一位
				helper[i][j] = helper[i-1][j-1]
			} else {
				/*否则等于
				i-1到j的操作加上一次插入
				i到j-1的操作加上一次删除
				i-1到j-1的操作加上一次修改
				三种情况的最小值
				*/
				helper[i][j] = min(
					helper[i-1][j]+1,
					helper[i][j-1]+1,
					helper[i-1][j-1]+1,
				)
			}
		}
	}
	return helper[m][n]
}

func min(a, b, c int) int {
	minTwo := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	return minTwo(minTwo(a, b), c)
}

// @lc code=end

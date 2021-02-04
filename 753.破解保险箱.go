/*
 * @lc app=leetcode.cn id=753 lang=golang
 *
 * [753] 破解保险箱
 */

// @lc code=start
package main

import (
	"bytes"
	"math"
)

type void struct{}

var (
	null   void
	set    map[int]void
	height int
)

func crackSafe(n int, k int) string {
	height = int(math.Pow10(n - 1))
	set = make(map[int]void)
	var result bytes.Buffer
	dfs(0, k, &result)
	result.Write(bytes.Repeat([]byte{'0'}, n-1))
	return result.String()
}

func dfs(node, k int, result *bytes.Buffer) {
	var neighbor int
	for i := 0; i < k; i++ {
		neighbor = node*10 + i
		if _, ok := set[neighbor]; !ok {
			set[neighbor] = null
			dfs(neighbor%height, k, result)
			(*result).WriteByte(byte(i + '0'))
		}
	}
}

// @lc code=end

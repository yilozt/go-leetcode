/*
 * @lc app=leetcode.cn id=547 lang=golang
 *
 * [547] 省份数量
 */

package main

// @lc code=start

func dfs(isConnected [][]int, visited []bool, city int) {
	if visited[city] == true {
		return
	}

	visited[city] = true
	for next_city, connected := range isConnected[city] {
		if connected == 1 {
			dfs(isConnected, visited, next_city)
		}
	}
}

func findCircleNum(isConnected [][]int) int {
	visited := make([]bool, len(isConnected))
	ans := 0

	for i := 0; i < len(isConnected); i++ {
		if !visited[i] {
			ans++
			dfs(isConnected, visited, i)
		}
	}

	return ans
}

// @lc code=end

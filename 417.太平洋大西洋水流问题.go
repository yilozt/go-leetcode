/*
 * @lc app=leetcode.cn id=417 lang=golang
 *
 * [417] 太平洋大西洋水流问题
 */

package main

import (
	"fmt"
	"math"
)

// @lc code=start
func dfs(heights, dirs [][]int, visited [][]bool, last_height, col int, row int, to_pacific, to_atlantic *bool) {
	if col < 0 || row < 0 || row >= len(heights) || col >= len(heights[0]) || visited[row][col] || heights[row][col] > last_height {
		return
	}
	if col == 0 || row == 0 {
		*to_pacific = true
	}
	if col == len(heights[0])-1 || row == len(heights)-1 {
		*to_atlantic = true
	}

	visited[row][col] = true
	for _, dir := range dirs {
		dfs(heights, dirs, visited, heights[row][col], col+dir[0], row+dir[1], to_pacific, to_atlantic)
		if *to_atlantic && *to_pacific {
			visited[row][col] = false
			return
		}
	}
	visited[row][col] = false
}

func pacificAtlantic(heights [][]int) [][]int {
	ans := make([][]int, 0)
	visited := make([][]bool, len(heights))
	for i := 0; i < len(heights); i++ {
		visited[i] = make([]bool, len(heights[0]))
	}
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for col := 0; col < len(heights[0]); col++ {
		for row := 0; row < len(heights); row++ {
			to_pacific, to_atlantic := false, false
			dfs(heights, dirs, visited, math.MaxInt, col, row, &to_pacific, &to_atlantic)
			if to_pacific && to_atlantic {
				ans = append(ans, []int{row, col})
			}
		}
	}
	return ans
}

// @lc code=end
func main() {
	fmt.Println(pacificAtlantic([][]int{{1, 2, 2, 3, 5}, {3, 2, 3, 4, 4}, {2, 4, 5, 3, 1}, {6, 7, 1, 4, 5}, {5, 1, 1, 2, 4}}))
}

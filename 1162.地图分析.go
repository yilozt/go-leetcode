/*
 * @lc app=leetcode.cn id=1162 lang=golang
 *
 * [1162] 地图分析
 */

package main

import "fmt"

// @lc code=start

type Queue []Pos

func (this *Queue) PushBack(a Pos) { *this = append(*this, a) }
func (this *Queue) PopFront() Pos  { a := (*this)[0]; *this = (*this)[1:]; return a }

type Pos struct {
	col int
	row int
}

func dfs(grid [][]int, col, row int) int {
	if grid[row][col] == 1 {
		return -1
	}

	queue := Queue{}
	queue.PushBack(Pos{col: col, row: row})

	visited := make([][]bool, 0)
	for i := 0; i < len(grid); i++ {
		visited = append(visited, make([]bool, len(grid[0])))
	}
	visited[row][col] = true

	dirs := [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	dist := 0
	for len(queue) != 0 {
		queue_len := len(queue)
		for i := 0; i < queue_len; i++ {
			cur := queue.PopFront()
			for _, dir := range dirs {
				nc := cur.col + dir[0]
				nr := cur.row + dir[1]

				if nc < 0 || nr < 0 || nc >= len(grid[0]) || nr >= len(grid) || visited[nr][nc] {
					continue
				}
				if grid[nr][nc] == 1 {
					return dist + 1
				}
				visited[nr][nc] = true
				queue.PushBack(Pos{col: nc, row: nr})
			}

		}
		dist++
	}

	return 0
}

func maxDistance(grid [][]int) int {
	if len(grid) == 0 {
		return -1
	}
	h := len(grid)
	w := len(grid[0])

	ans := -1
	for col := 0; col < w; col++ {
		for row := 0; row < h; row++ {
			dist := dfs(grid, col, row)
			if dist == 0 {
				return -1
			}
			if ans < dist {
				ans = dist
			}
		}
	}
	return ans
}

// @lc code=end

func main() {
	desc := []struct {
		grid [][]int
		res  int
	}{
		{grid: [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}, res: 2},
		{grid: [][]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 0}}, res: 4},
		{grid: [][]int{{0, 0, 0}, {0, 0, 0}}, res: -1},
		{grid: [][]int{{1, 1, 1}, {1, 1, 1}}, res: -1},
	}

	for _, des := range desc {
		got := maxDistance(des.grid)
		want := des.res

		if got != want {
			fmt.Printf("got: %d, want: %d, grid: %v\n", got, want, des.grid)
		}
		fmt.Printf("==============\n")
	}
}

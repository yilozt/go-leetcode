/*
 * @lc app=leetcode.cn id=1020 lang=golang
 *
 * [1020] 飞地的数量
 */

package main

import "fmt"

// @lc code=start

type Queue []int

func (q *Queue) PushBack(col int, row int)    { *q = append(*q, col, row) }
func (q *Queue) PopFront() (col int, row int) { col = (*q)[0]; row = (*q)[1]; *q = (*q)[2:]; return }

// 尝试遍历 (col, row) 处的岛屿，返回岛屿的面积
// 在遍历时将岛屿淹没
// 如果岛屿与地图边界相交，则返回 0
func travel(grid [][]int, col int, row int) int {
	if grid[row][col] == 0 {
		return 0
	}

	// 地图尺寸
	w := len(grid[0])
	h := len(grid)

	// 广度优先遍历 - 使用队列缓存节点
	q := Queue{}
	q.PushBack(col, row)

	// 岛屿面积
	area := 0

	// 是否与地图边界相交
	crossBound := false

	for len(q) != 0 {
		cur_col, cur_row := q.PopFront()

		if cur_col < 0 || cur_row < 0 || cur_col >= w || cur_row >= h {
			crossBound = true
			continue
		}

		if grid[cur_row][cur_col] == 0 {
			continue
		}

		// 将四周的单元格添加到队列
		q.PushBack(cur_col+0, cur_row+1)
		q.PushBack(cur_col+0, cur_row-1)
		q.PushBack(cur_col-1, cur_row+0)
		q.PushBack(cur_col+1, cur_row+0)

		// 找到一个陆地，将其淹没
		// 面积 + 1
		grid[cur_row][cur_col] = 0
		area += 1
	}

	if crossBound {
		return 0
	} else {
		return area
	}
}

func numEnclaves(grid [][]int) int {
	count := 0

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			count += travel(grid, col, row)
		}
	}

	return count
}

// @lc code=end

func main() {
	grid := [][]int{{0, 0, 0, 0}, {1, 0, 1, 0}, {0, 1, 1, 0}, {0, 0, 0, 0}}
	fmt.Println(numEnclaves(grid))

	grid = [][]int{{0, 1, 1, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 0, 0}}
	fmt.Println(numEnclaves(grid))
}

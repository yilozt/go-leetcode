/*
 * @lc app=leetcode.cn id=1905 lang=golang
 *
 * [1905] 统计子岛屿
 */

package main

// @lc code=start
type Queue []int

func (q *Queue) PushBack(col int, row int)    { *q = append(*q, col, row) }
func (q *Queue) PopFront() (col int, row int) { col = (*q)[0]; row = (*q)[1]; *q = (*q)[2:]; return }

// 在遍历 grid2 的岛屿时，查看 grid1 相同位置情况，如果这个区域也是陆地
// 那么 grid2 的岛屿是子岛屿
func travel(grid1 [][]int, grid2 [][]int, col int, row int) int {
	if grid2[row][col] == 0 {
		return 0
	}

	// 地图尺寸
	w := len(grid1[0])
	h := len(grid1)

	// 广度优先遍历 - 使用队列缓存节点
	q := Queue{}
	q.PushBack(col, row)

	fail := false

	for len(q) != 0 {
		cur_col, cur_row := q.PopFront()

		if cur_col < 0 || cur_row < 0 || cur_col >= w || cur_row >= h || grid2[cur_row][cur_col] == 0 {
			continue
		}

		// 将四周的单元格添加到队列
		q.PushBack(cur_col+0, cur_row+1)
		q.PushBack(cur_col+0, cur_row-1)
		q.PushBack(cur_col-1, cur_row+0)
		q.PushBack(cur_col+1, cur_row+0)

		// 看 grid1 同位置是否为陆地
		if grid1[cur_row][cur_col] == 0 {
			fail = true
		}

		// 找到一个陆地
		grid2[cur_row][cur_col] = 0
	}

	if fail {
		return 0
	} else {
		return 1
	}
}

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	count := 0

	for row := 0; row < len(grid1); row++ {
		for col := 0; col < len(grid1[0]); col++ {
			count += travel(grid1, grid2, col, row)
		}
	}

	return count
}

// @lc code=end

/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */

package main

// @lc code=start
import "fmt"

type Queue ([]int)

func (s *Queue) PushBack(i int) { *s = append(*s, i) }
func (s *Queue) PopFront() int  { head := (*s)[0]; *s = (*s)[1:]; return head }

func numIslands(grid [][]byte) int {
	fmt.Println(grid)

	lc, lr := len(grid[0]), len(grid)
	count := 0

	for c := 0; c < lc; c++ {
		for r := 0; r < lr; r++ {
			if FillLand(&grid, r, c) {
				count++
			}
		}
	}

	return count
}

func FillLand(grid *[][]byte, r int, c int) bool {
	if (*grid)[r][c] != 49 {
		return false
	}

	lc, lr := len((*grid)[0]), len((*grid))

	try := func(q *Queue, color byte, cr int, cc int) {
		if cc < 0 || cr < 0 || cc >= lc || cr >= lr {
			return
		}
		n_color := (*grid)[cr][cc]
		if n_color == color && n_color != 2 {
			q.PushBack(cr)
			q.PushBack(cc)
		}
	}

	q := Queue([]int{})

	q.PushBack(r)
	q.PushBack(c)

	for len(q) != 0 {
		cr := q.PopFront()
		cc := q.PopFront()

		color := (*grid)[cr][cc]
		try(&q, color, cr+0, cc+1)
		try(&q, color, cr+0, cc-1)
		try(&q, color, cr+1, cc+0)
		try(&q, color, cr-1, cc+0)

		(*grid)[cr][cc] = 2
	}

	return true
}

// @lc code=end

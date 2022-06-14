/*
 * @lc app=leetcode.cn id=733 lang=golang
 *
 * [733] 图像渲染
 */

package main

import "fmt"

// @lc code=start

type Pos struct {
	sc int
	sr int
}
type Queue ([]Pos)

func (q *Queue) PopFront() Pos  { h := (*q)[0]; *q = (*q)[1:]; return h }
func (q *Queue) PushBack(i Pos) { *q = append(*q, i) }

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	q := Queue([]Pos{})
	lc, lr := len(image[0]), len(image)

	q.PushBack(Pos{sc: sc, sr: sr})

	dirs := [...]Pos{{sc: 0, sr: 1}, {sc: 0, sr: -1}, {sc: 1, sr: 0}, {sc: -1, sr: 0}}

	for len(q) != 0 {
		cur := q.PopFront()
		old_color := &image[cur.sr][cur.sc]

		for _, dir := range dirs {
			next := Pos{sc: cur.sc + dir.sc, sr: cur.sr + dir.sr}
			if next.sc < 0 || next.sc >= lc || next.sr < 0 || next.sr >= lr {
				continue
			}

			new_color := image[next.sr][next.sc]
			if new_color != newColor && new_color == *old_color {
				q.PushBack(next)
			}
		}

		*old_color = newColor

	}

	return image
}

// @lc code=end

func main() {
	q := Queue([]Pos{})
	q.PushBack(Pos{})
	fmt.Println("Len: ", len(q))
}

/*
 * @lc app=leetcode.cn id=59 lang=golang
 *
 * [59] 螺旋矩阵 II
 */

package main

import "fmt"

// @lc code=start
type Pos struct {
	col int
	row int
}

type Bound struct {
	leftTop     Pos
	rightBottom Pos
}

func (this *Bound) empty() bool {
	return this.leftTop.col > this.rightBottom.col || this.leftTop.row > this.rightBottom.row
}
func (this *Bound) shrinkTop()    { this.leftTop.row++ }
func (this *Bound) shrinkBottom() { this.rightBottom.row-- }
func (this *Bound) shrinkLeft()   { this.leftTop.col++ }
func (this *Bound) shrinkRight()  { this.rightBottom.col-- }
func (this *Bound) OutOfBound(col int, row int) bool {
	return col < this.leftTop.col || row < this.leftTop.row || col > this.rightBottom.col || row > this.rightBottom.row
}

func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	dirs := []Pos{{col: 1, row: 0}, {col: 0, row: 1}, {col: -1, row: 0}, {col: 0, row: -1}}
	dir := 0

	bound := Bound{
		leftTop:     Pos{col: 0, row: 0},
		rightBottom: Pos{col: n - 1, row: n - 1},
	}

	count := 1
	cur := Pos{col: 0, row: 0}

	for !bound.empty() {
		if bound.OutOfBound(cur.col, cur.row) {
			// 恢复位置
			d := dirs[dir]
			cur.col -= d.col
			cur.row -= d.row

			switch dir {
			case 0:
				bound.shrinkTop()
			case 1:
				bound.shrinkRight()
			case 2:
				bound.shrinkBottom()
			case 3:
				bound.shrinkLeft()
			default:
			}
			dir = (dir + 1) % 4
		} else {
			res[cur.row][cur.col] = count
			count++
		}

		d := dirs[dir]
		cur.col += d.col
		cur.row += d.row
	}

	return res
}

// @lc code=end

func main() {
	fmt.Println(generateMatrix(3))
	fmt.Println(generateMatrix(1))
	fmt.Println(generateMatrix(4))
}

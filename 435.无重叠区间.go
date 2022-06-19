/*
 * @lc app=leetcode.cn id=435 lang=golang
 *
 * [435] 无重叠区间
 */

package main

import (
	"fmt"
	"sort"
)

// @lc code=start
func eraseOverlapIntervals(intervals [][]int) int {
	len := len(intervals)
	if len == 0 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool { return intervals[i][1] < intervals[j][1] })

	ans := 1
	right := intervals[0][1]
	for i := 1; i < len; i++ {
		if intervals[i][0] >= right {
			ans++
			right = intervals[i][1]
		}
	}

	return len - ans
}

// @lc code=end

func main() {
	desc := []struct {
		inputs [][]int
		output int
	}{
		{inputs: [][]int{{1, 100}, {11, 22}, {1, 11}, {2, 12}}, output: 2},
		{inputs: [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}, output: 1},
		{inputs: [][]int{{1, 2}, {1, 2}, {1, 2}}, output: 2},
		{inputs: [][]int{{1, 2}, {2, 3}}, output: 0},
		{inputs: [][]int{{1, 2}}, output: 0},
		{inputs: [][]int{}, output: 0},
	}

	for _, dec := range desc {
		got := eraseOverlapIntervals(dec.inputs)
		want := dec.output
		if got != want {
			fmt.Printf("FAILED: got: %d, want: %d, input: %v\n", got, want, dec.inputs)
		}
	}
}

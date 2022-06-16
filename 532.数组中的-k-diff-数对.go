/*
 * @lc app=leetcode.cn id=532 lang=golang
 *
 * [532] 数组中的 k-diff 数对
 */

package main

// @lc code=start

type Void struct{}

var Empty = Void{}

func findPairs(nums []int, k int) int {
	visited, res := make(map[int]Void, 0), make(map[int]Void, 0)

	for _, num := range nums {
		_, ok := visited[num-k]
		if ok {
			res[num-k] = Empty
		}
		_, ok = visited[num+k]
		if ok {
			res[num] = Empty
		}
		visited[num] = Empty
	}

	return len(res)
}

// @lc code=end

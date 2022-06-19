/*
 * @lc app=leetcode.cn id=986 lang=golang
 *
 * [986] 区间列表的交集
 */

package main

// @lc code=start
func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	ptr_f := 0
	ptr_s := 0
	ans := make([][]int, 0)

	for ptr_f < len(firstList) && ptr_s < len(secondList) {
		first := firstList[ptr_f]
		second := secondList[ptr_s]

		// 如果相交
		if !(first[1] < second[0] || second[1] < first[0]) {
			ans = append(ans, []int{max(first[0], second[0]), min(first[1], second[1])})
		}

		if first[1] == second[1] {
			ptr_f++
			ptr_s++
		} else if first[1] < second[1] {
			ptr_f++
		} else {
			ptr_s++
		}
	}

	return ans
}

func max(i int, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}
func min(i int, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}

// @lc code=end

/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

package main

// @lc code=start
import "sort"

type void struct{}

var empty void

func threeSum(nums []int) [][]int {
	len := len(nums)

	if len < 3 {
		return [][]int{}
	}

	hash_map := make(map[int]void, 0)
	res := make([][]int, 0)

	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })

	for _, n := range nums {
		hash_map[n] = empty
	}

	i := 0
	for i < len-2 {
		j := i + 1
		for j < len-1 {
			rest := -(nums[i] + nums[j])

			if rest < nums[j+1] {
				break
			}

			_, ok := hash_map[rest]
			if ok {
				res = append(res, []int{nums[i], nums[j], rest})
			}

			// skip repeat elem
			j++
			for j < len && nums[j] == nums[j-1] {
				j++
			}
		}

		// skip repeat elem
		i++
		for i < len && nums[i] == nums[i-1] {
			i++
		}
	}

	return res
}

// @lc code=end

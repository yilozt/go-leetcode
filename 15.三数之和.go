/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int
	for i := 0; i < len(nums); i++ {
		// 跳过重复元素
		if i == 0 || nums[i] != nums[i-1] {
			target := -nums[i]
			for j := i + 1; j < len(nums); j++ {
				if j == i+1 || nums[j] != nums[j-1] {
					k := len(nums) - 1
					for nums[j]+nums[k] > target && k > j {
						k--
					}
					if j == k {
						break
					}
					if nums[j]+nums[k] == target {
						result = append(result, []int{nums[i], nums[j], nums[k]})
					}
				}
			}
		}
	}
	return result
}

// @lc code=end

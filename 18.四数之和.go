/*
 * @lc app=leetcode.cn id=18 lang=golang
 *
 * [18] 四数之和
 */

// @lc code=start
func fourSum(nums []int, target int) [][]int {
	var result [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums)-3; i++ {
		if i == 0 || nums[i] != nums[i-1] {
			if nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target {
				break
			}
			if nums[i]+nums[len(nums)-1]+nums[len(nums)-2]+nums[len(nums)-3] < target {
				continue
			}
			for j := i + 1; j < len(nums)-2; j++ {
				if j == i+1 || nums[j] != nums[j-1] {
					if nums[i]+nums[j]+nums[j+1]+nums[j+2] > target {
						break
					}
					if nums[i]+nums[j]+nums[len(nums)-1]+nums[len(nums)-2] < target {
						continue
					}
					k, l := j+1, len(nums)-1
					for k < l {
						tmp := nums[i] + nums[j] + nums[k] + nums[l]
						if tmp <= target {
							if tmp == target {
								result = append(result, []int{nums[i], nums[j], nums[k], nums[l]})
							}
							for k < l && nums[k] == nums[k+1] {
								k++
							}
							k++
						} else if tmp > target {
							for k < l && nums[l] == nums[l-1] {
								l--
							}
							l--
						}
					}
				}
			}
		}
	}
	return result
}

// @lc code=end

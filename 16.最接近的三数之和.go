/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 */

// @lc code=start
func threeSumClosest(nums []int, target int) int {
	// 最终结果
	result := 0
	// 与target的距离
	distance := math.MaxInt32
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i] != nums[i-1] {
			j := i + 1
			k := len(nums) - 1
			for j < k {
				tmp := nums[i] + nums[j] + nums[k]
				dis := tmp - target
				if dis > 0 {
					if dis < distance {
						result = tmp
						distance = dis
					}
					for j < k && nums[k] == nums[k-1] {
						k--
					}
					k--
				} else if dis < 0 {
					if (-dis) < distance {
						result = tmp
						distance = (-dis)
					}
					for j < k && nums[j] == nums[j+1] {
						j++
					}
					j++
				} else {
					return target
				}
			}
		}
	}
	return result
}

// @lc code=end

/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] 下一个排列
 */

// @lc code=start
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}
	var small, big int
	flag := true
	for small = len(nums) - 2; small >= 0; small-- {
		if nums[small] < nums[small+1] {
			flag = false
			break
		}
	}
	if flag {
		reverse(nums)
		return
	}
	for big = len(nums) - 1; big > small; big-- {
		if nums[small] < nums[big] {
			nums[small], nums[big] = nums[big], nums[small]
			break
		}
	}
	reverse(nums[small+1:])
}

// @lc code=end

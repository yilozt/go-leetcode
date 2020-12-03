/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	result := 0
	var h int
	for left < right {
		width := right - left
		if height[left] >= height[right] {
			h = height[right]
			right -= 1
		} else {
			h = height[left]
			left += 1
		}
		result = max(h*width, result)
	}
	return result
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end


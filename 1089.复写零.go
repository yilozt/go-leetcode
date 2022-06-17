/*
 * @lc app=leetcode.cn id=1089 lang=golang
 *
 * [1089] 复写零
 */

package main

import "fmt"

// @lc code=start
func duplicateZeros(arr []int) {
	len := len(arr)

	if len < 1 {
		return
	}

	end := 0
	i := 0

	count_not_zero := 0
	count_zero := 0

	for {
		if arr[end] == 0 {
			i += 2
			count_zero += 1
		} else {
			i += 1
			count_not_zero += 1
		}
		if i >= len {
			break
		}
		end++
	}

	i = len - 1

	// 考虑最后一个0被截断的情况
	if count_zero*2+count_not_zero > len {
		arr[i] = 0
		i--
		end--
	}

	// 从后往前填写数组
	for i >= 0 && end >= 0 {
		// 重写 0
		if arr[end] == 0 {
			arr[i] = 0
			i--
			arr[i] = 0
		} else {
			arr[i] = arr[end]
		}
		i--
		end--
	}

}

// @lc code=end
func main() {
	duplicateZeros([]int{0, 0, 0, 0})
	duplicateZeros([]int{0, 0, 0})

	a := []int{8, 4, 5, 0, 0, 0, 0, 7}
	duplicateZeros(a)
	fmt.Println(a)
}

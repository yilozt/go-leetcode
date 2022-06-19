/*
 * @lc app=leetcode.cn id=415 lang=golang
 *
 * [415] 字符串相加
 */

package main

import "fmt"

// @lc code=start
func addStrings(num1 string, num2 string) string {
	ans := ""

	len_num1, len_num2 := len(num1), len(num2)
	len := len_num1
	if len < len_num2 {
		len = len_num2
	}

	overflow := 0
	for i := 0; i < len; i++ {
		n1 := 0
		n2 := 0
		if len_num1-1-i >= 0 {
			n1 = int(num1[len_num1-1-i] - '0')
		}
		if len_num2-1-i >= 0 {
			n2 = int(num2[len_num2-1-i] - '0')
		}

		res := n1 + n2 + overflow

		if res >= 10 {
			overflow = 1
			res -= 10
		} else {
			overflow = 0
		}
		ans = fmt.Sprint(res) + ans
	}
	if overflow != 0 {
		ans = fmt.Sprint(overflow) + ans
	}
	return ans
}

// @lc code=end
func main() {
	fmt.Println(addStrings("456", "77"))
	fmt.Println(addStrings("123456789", "987654321"))
}

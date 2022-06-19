/*
 * @lc app=leetcode.cn id=844 lang=golang
 *
 * [844] 比较含退格的字符串
 */

package main

import "fmt"

// @lc code=start
func processBackspace(s *string) {
	len := 0
	bytes := []byte(*s)
	for _, char := range bytes {
		if char == '#' {
			len--
			if len < 0 {
				len = 0
			}
		} else {
			bytes[len] = byte(char)
			len++
		}
	}
	*s = string(bytes[0:len])
}

func backspaceCompare(s string, t string) bool {
	processBackspace(&s)
	processBackspace(&t)
	return s == t
}

// @lc code=end

func main() {
	s := "####aabc#cc"
	processBackspace(&s)
	fmt.Println(string(s))
}

/*
 * @lc app=leetcode.cn id=394 lang=golang
 *
 * [394] 字符串解码
 */

// @lc code=start
package main

import (
	"strconv"
	"strings"
	"unicode"
)

func decodeString(s string) string {
	var (
		stack            []rune
		numBegin, numEnd int
		strBegin, strEnd int
	)
	for _, item := range s {
		if item == ']' {
			strEnd, strBegin = len(stack), len(stack)-1
			for stack[strBegin] != '[' {
				strBegin--
			}
			str := string(stack[strBegin+1 : strEnd])
			numBegin, numEnd = strBegin-1, strBegin
			for numBegin >= 0 && unicode.IsNumber(stack[numBegin]) {
				numBegin--
			}
			num, _ := strconv.Atoi(string(stack[numBegin+1 : numEnd]))
			tmp := strings.Repeat(str, num)
			stack = stack[:numBegin+1]
			stack = []rune(string(stack) + tmp)
		} else {
			stack = append(stack, item)
		}
	}
	return string(stack)
}

// @lc code=end

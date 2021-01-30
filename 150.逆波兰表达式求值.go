/*
 * @lc app=leetcode.cn id=150 lang=golang
 *
 * [150] 逆波兰表达式求值
 */

// @lc code=start
package main

import "strconv"

func evalRPN(tokens []string) int {
	var stack []int
	for _, item := range tokens {
		switch item {
		case "+", "-", "*", "/":
			first, second := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			var tmp int
			switch item {
			case "+":
				tmp = first + second
			case "-":
				tmp = first - second
			case "*":
				tmp = first * second
			case "/":
				tmp = first / second
			}
			stack = append(stack, tmp)
		default:
			value, _ := strconv.Atoi(item)
			stack = append(stack, value)
		}
	}
	return stack[0]
}

// @lc code=end

package main

import "strconv"

func findNthDigit(n int) int {
	digit, start, count := 1, 1, 9
	for n > count {
		n -= count
		start *= 10
		digit += 1
		count = 9 * start * digit
	}
	// 当前的数字
	currentNum := start + (n-1)/digit
	return int(strconv.Itoa(currentNum)[(n-1)%digit] - '0')
}

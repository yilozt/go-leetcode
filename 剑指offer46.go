package main

import (
	"strconv"
)

func translateNum(num int) int {
	str := strconv.Itoa(num)
	f := make([]int, len(str))
	for i := len(str) - 1; i >= 0; i-- {
		count := 0 // 从i开始的翻译的种数
		if i < len(str)-1 {
			count = f[i+1] // 首先继承从i+1开始翻译的种数
			if result, err := strconv.Atoi(str[i : i+2]); err == nil {
				if result >= 10 && result <= 25 {
					if i < len(str)-2 {
						count += f[i+2]
					} else {
						count += 1
					}
				}
			}
		} else {
			count = 1 // 最后一位，只有一种翻译方式
		}
		f[i] = count
	}
	return f[0]
}

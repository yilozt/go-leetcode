/*
题目描述（使用OCR生成）：
小明和小红在做游戏,由小明先抽到一串小写字母的随机字符申,小红也抽一个比小明更长的
字符串,也由小写字母组成,并且小红可以有限次的替换一个字母为另一个字母。如:
小明的串是abba;
小红的串是adbba;
最大替换次数2.
那么小红可以通过把d修改为a,从而使小明的串变成小红的子字符串,而且没有超过最大替
换次数2,这样就成功了。
现在请你帮忙判断:小红每次游戏时能不能成功替换字母从而包含小明的字串呢,如果可
以,输出最少替换次数,如果不可以或者不需要替换就满足,输出0。
输入描述：

*/
package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		first, second string
		maxCount      int
	)
	fmt.Scanln(&first)
	fmt.Scanln(&second)
	fmt.Scanln(&maxCount)
	// 暴力解法
	result := math.MaxInt32
	for i := 0; i <= len(second)-len(first); i++ {
		var tmp int
		for j := 0; j < len(first); j++ {
			if first[j] != second[i+j] {
				tmp += 1
			}
		}
		result = min(result, tmp)
	}
	if result > maxCount {
		fmt.Println(0)
	}
	fmt.Println(result)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

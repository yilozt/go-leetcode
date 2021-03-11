package main

import "fmt"

func main() {
	var (
		n, m  int
		input string
	)
	fmt.Scanln(&n, &m)
	fmt.Scanln(&input)
	// 左右指针
	left, right := 0, 0
	// 最后结果
	result := -1
	// 统计字母在窗口中的出现次数
	var count [2]int
	for right < n {
		char := input[right]
		right++
		count[int(char-'a')] += 1
		// 加入该字母后操作次数超过上限
		if min(count[0], count[1]) == m+1 {
			// 更新最大长度
			result = max(result, right-left-1)
			for min(count[0], count[1]) == m+1 {
				char = input[left]
				left++
				count[int(char-'a')]--
			}
		}
	}
	if result == -1 {
		result = n
	}
	fmt.Println(result)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

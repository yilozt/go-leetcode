package main

import (
	"sort"
)

func permutation(s string) []string {
	// 保存最后结果
	var result []string
	// 字符串
	str := []byte(s)
	// 排序输入的字符串
	sort.Slice(str, func(i, j int) bool {
		return str[i] < str[j]
	})
	visited := make([]bool, len(str))
	// 使用回溯方法
	var backTrack func([]byte)
	// 闭包函数
	backTrack = func(current []byte) {
		// 符合长度条件，添加到解集合中
		if len(current) == len(str) {
			result = append(result, string(current))
			return
		}
		for i := 0; i < len(str); i++ {
			if visited[i] {
				continue
			}
			// 这个判断保证重复字符必定按照从左到右的顺序填入
			// 避免了重复字符排列的问题
			if i != 0 && str[i] == str[i-1] && !visited[i-1] {
				continue
			}
			current = append(current, str[i])
			visited[i] = true
			backTrack(current)
			visited[i] = false
			current = current[:len(current)-1]
		}
	}
	backTrack([]byte{})
	return result
}

package main

import (
	"sort"
	"strconv"
	"strings"
)

type stringList []string

func (s stringList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s stringList) Len() int {
	return len(s)
}

func (s stringList) Less(i, j int) bool {
	return s[i]+s[j] < s[j]+s[i]
}

func minNumber(nums []int) string {
	var strs stringList
	// 将数字转换为字符串
	for _, v := range nums {
		strs = append(strs, strconv.Itoa(v))
	}
	sort.Sort(strs)
	return strings.Join(strs, "")
}

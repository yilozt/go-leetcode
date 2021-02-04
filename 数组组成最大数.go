package main

import (
	"fmt"
	"sort"
	"strings"
)

type strs []string

func (s strs) Len() int {
	return len(s)
}

func (s strs) Less(i, j int) bool {
	return s[i]+s[j] > s[j]+s[i]
}

func (s strs) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	var input string
	for {
		_, err := fmt.Scan(&input)
		if err != nil {
			return
		}
		input = strings.Replace(strings.Replace(input, "[", "", 1), "]", "", 1)
		inputs := strs(strings.Split(input, ","))
		sort.Sort(inputs)
		resultStr := strings.Join(inputs, "")
		fmt.Println(resultStr)
	}
}

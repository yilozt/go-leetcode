package main

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
)

// 判断是否是数字
func isNum(b byte) bool {
	if b >= '0' && b <= '9' {
		return true
	}
	return false
}

// 判断是否是大写字母
func isUppercase(b byte) bool {
	if b >= 'A' && b <= 'Z' {
		return true
	}
	return false
}

// 判断是否是小写字母
func isLowercase(b byte) bool {
	if b >= 'a' && b <= 'z' {
		return true
	}
	return false
}

func main() {
	var input string
	for {
		_, err := fmt.Scan(&input)
		if err != nil {
			return
		}
		fmt.Println(parse(input))
	}
}

func parse(input string) string {
	var (
		stack, token         []byte
		tokenEnd, tokenBegin int
		countBegin, countEnd int
		count                int
		label                byte
	)
	for i := 0; i < len(input); i++ {
		switch input[i] {
		case ')', ']':
			// 得到括号内的内容作为token
			tokenEnd, tokenBegin = len(stack), len(stack)-1
			switch input[i] {
			case ')':
				label = '('
			case ']':
				label = '['
			}
			for stack[tokenBegin] != label {
				tokenBegin--
			}
			// 取到token
			token = stack[tokenBegin+1 : tokenEnd]
			// 栈弹出到label之前
			stack = stack[:tokenBegin]
			// 取到括号后跟的数字作为count
			countBegin, countEnd = i+1, i+1
			for countEnd < len(input) && isNum(input[countEnd]) {
				countEnd++
			}
			// 取到count
			tmpCount := input[countBegin:countEnd]
			if len(tmpCount) == 0 {
				count = 1
			} else {
				count, _ = strconv.Atoi(input[countBegin:countEnd])
			}
			// 将token和count传给handle函数做进一步统计，返回统计好的字符串
			result := handleTokenAndCount(token, count, false)
			// 将统计好的字符串压入栈中
			stack = append(stack, result...)
			i = countEnd - 1
		default:
			stack = append(stack, input[i])
		}
	}
	// 这样处理过后，stack内是不含括号的元素个数统计,最后需要对其排序并输出
	return handleTokenAndCount(stack, 1, true)
}

func handleTokenAndCount(token []byte, count int, order bool) string {
	var tmp bytes.Buffer
	var element, tmpCount string
	result := make(map[string]int)
	// 首先统计token中的元素种类和个数
	for i := 0; i < len(token); i++ {
		if isUppercase(token[i]) {
			// tmp存储元素名称
			tmp.Reset()
			tmp.WriteByte(token[i])
			for i+1 < len(token) && isLowercase(token[i+1]) {
				tmp.WriteByte(token[i+1])
				i++
			}
			// 存储该元素名称
			element = tmp.String()
			// tmp存储出现次数
			tmp.Reset()
			for i+1 < len(token) && isNum(token[i+1]) {
				tmp.WriteByte(token[i+1])
				i++
			}
			tmpCount = tmp.String()
			// 存储元素出现次数
			if len(tmpCount) == 0 {
				result[element] += 1
			} else {
				tmpCountInt, _ := strconv.Atoi(tmpCount)
				result[element] += tmpCountInt
			}
		}
	}
	// 接着将统计完成的元素个数翻count倍
	for i, v := range result {
		result[i] = v * count
	}
	tmp.Reset()
	// 根据参数返回排序/乱序的结果
	if order {
		var helper []string
		// 保存key
		for i := range result {
			helper = append(helper, i)
		}
		sort.Strings(helper)
		for _, v := range helper {
			tmp.WriteString(v)
			tmp.WriteString(strconv.Itoa(result[v]))
		}
	} else {
		for i, v := range result {
			tmp.WriteString(i)
			tmp.WriteString(strconv.Itoa(v))
		}
	}
	return tmp.String()
}

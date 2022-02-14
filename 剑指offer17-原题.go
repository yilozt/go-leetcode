package main

import "fmt"

func main() {
	printNumbers(13)
}

func printNumbers(n int) { // 注意n可能过大导致类型溢出，因此使用字符串方式处理
	numbers := make([]byte, n)
	innerPrintNumbers(numbers, n, 0)
}

func innerPrintNumbers(numbers []byte, n, current int) {
	if current == n {
		fmt.Println(string(numbers))
		return
	}
	var i byte
	for i = 0; i < 10; i++ {
		numbers[current] = '0' + i
		innerPrintNumbers(numbers, n, current+1)
	}
}

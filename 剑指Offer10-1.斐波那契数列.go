package main

import "fmt"

func f(n int, tmp []int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	if tmp[n] != 0 {
		return tmp[n]
	}

	tmp[n] = (f(n-1, tmp) + f(n-2, tmp)) % (1e9 + 7)
	return tmp[n]
}

func fib(n int) int {
	tmp := make([]int, 102)
	res := f(n, tmp)
	return res
}

func main() {
	fmt.Println("flib(2) : ", fib(2))
	fmt.Println("flib(5) : ", fib(5))
	fmt.Println("flib(45) : ", fib(45))
}

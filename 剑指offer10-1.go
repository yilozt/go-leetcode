package main

func fib(n int) int {
	if n < 2 {
		return n
	}
	current := 0
	currentNext := 1
	for i := 2; i <= n; i++ {
		current, currentNext = currentNext, current+currentNext
	}
	return currentNext
}

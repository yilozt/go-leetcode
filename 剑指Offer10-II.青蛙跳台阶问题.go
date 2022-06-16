package main

import "fmt"

// f(n) = f(n-1) + f(n-2)
// 使用循环
func numWays(n int) int {
	prev := 1
	cur := 2

	if n < 2 {
		return prev
	}

	for i := 0; i < n-2; i++ {
		tmp := cur
		cur += prev
		cur %= (1e9 + 7)
		prev = tmp
	}

	return cur
}

func main() {
	p := []struct {
		input  int
		output int
	}{
		{input: 2, output: 2}, {input: 7, output: 21}, {input: 0, output: 1},
	}

	for _, problem := range p {
		got := numWays(problem.input)
		want := problem.output

		if got != want {
			fmt.Printf("got: %v, want: %v", got, want)
		}
	}
}

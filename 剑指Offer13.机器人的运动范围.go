package main

import "fmt"

func movingCount(m int, n int, k int) int {
	visited := make([][]bool, 0)
	for i := 0; i < m; i++ {
		visited = append(visited, make([]bool, n))
	}
	count := 0
	dfs(m, n, k, 0, 0, visited, &count)
	return count
}

func sumOfnum(n int) (sum int) {
	for i := n; i != 0; i /= 10 {
		sum += i % 10
	}

	return
}

func dfs(m, n, k, col, row int, visited [][]bool, count *int) {
	if col < 0 || row < 0 || row >= m || col >= n ||
		visited[row][col] ||
		sumOfnum(col)+sumOfnum(row) > k {
		return
	}

	*count += 1

	visited[row][col] = true
	dfs(m, n, k, col+1, row+0, visited, count)
	dfs(m, n, k, col-1, row+0, visited, count)
	dfs(m, n, k, col+0, row+1, visited, count)
	dfs(m, n, k, col+0, row-1, visited, count)
}

func main() {
	fmt.Printf("got: %v, want : %v\n", movingCount(2, 3, 1), 3)
	fmt.Printf("got: %v, want : %v\n", movingCount(3, 1, 0), 1)
	fmt.Printf("got: %v, want : %v\n", movingCount(3, 2, 17), 6)
	fmt.Printf("got: %v, want : %v\n", movingCount(11, 8, 16), 88)
}

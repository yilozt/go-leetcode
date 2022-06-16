/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 */

package main

import "fmt"

// @lc code=start
func exist(board [][]byte, word string) bool {
	visited := make([][]bool, 0)
	for i := 0; i < len(board); i++ {
		visited = append(visited, make([]bool, len(board[0])))
	}
	res := false

	str := []byte(word)

	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[0]); col++ {
			res = res || dfs(board, visited, str, 0, col, row)
		}
	}

	return res
}

func dfs(board [][]byte, visited [][]bool, word []byte, depth int, col int, row int) bool {
	if depth == len(word) {
		return true
	}

	if col < 0 || row < 0 || col >= len(board[0]) || row >= len(board) || board[row][col] != word[depth] || visited[row][col] == true {
		return false
	}

	visited[row][col] = true

	res := dfs(board, visited, word, depth+1, col, row+1) ||
		dfs(board, visited, word, depth+1, col, row-1) ||
		dfs(board, visited, word, depth+1, col+1, row) ||
		dfs(board, visited, word, depth+1, col-1, row)

	visited[row][col] = false

	return res
}

// @lc code=end

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'},
	}

	tests := []struct {
		board  [][]byte
		word   string
		answer bool
	}{
		{board: board, word: "ABCCED", answer: true},
		{board: board, word: "SEE", answer: true},
		{board: board, word: "ABCB", answer: false},
	}

	for _, test := range tests {
		got := exist(test.board, test.word)
		fmt.Printf("err: got: %v, want: %v\n", got, test.answer)
	}
}

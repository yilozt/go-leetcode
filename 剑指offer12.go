package main

var directions = [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if search(board, i, j, word, 0) {
				return true
			}
		}
	}
	return false
}

// 参数说明：board 矩阵，i j 矩阵坐标，word 需构成的词 index 当前查找到词的位置
func search(board [][]byte, i, j int, word string, index int) bool {
	if index == len(word) { // 已经找到
		return true
	}
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) { // 查找到该字母时越界
		return false
	}
	if board[i][j] == word[index] { // 找到该字母
		tmp := board[i][j]
		board[i][j] = ' ' // 清空该位置，保证不重复使用
		for idx := range directions {
			if search(board, i+directions[idx][0], j+directions[idx][1], word, index+1) {
				return true // 在周边位置搜索下一个字母
			}
		}
		// 没搜到，回溯
		board[i][j] = tmp
	}
	return false
}

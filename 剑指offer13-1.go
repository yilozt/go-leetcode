package main

// 递归方法

var directions = [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func digitSum(item int) int {
	var sum int
	for item > 0 {
		sum += item % 10
		item /= 10
	}
	return sum
}

func movingCount(m int, n int, k int) int {
	visited := make([]bool, m*n)
	return movingCountCore(0, 0, m, n, k, visited)
}

func movingCountCore(currentx, currenty, m, n, k int, visited []bool) int {
	if currentx >= 0 && currentx < m && // x未越界
		currenty >= 0 && currenty < n && // y未越界
		digitSum(currentx)+digitSum(currenty) <= k && // 符合条件
		!visited[currentx*n+currenty] { // 未访问过
		visited[currentx*n+currenty] = true // 标记为访问
		result := 1
		for idx := range directions {
			result += movingCountCore(currentx+directions[idx][0],
				currenty+directions[idx][1],
				m, n, k, visited)
		}
		return result
	}
	return 0
}

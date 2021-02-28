package main

func getCount(m int) int {
	result := 0
	for m > 0 {
		result += m % 10
		m /= 10
	}
	return result
}

func movingCount(m int, n int, k int) int {
	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}
	result := 0
	vis[0][0] = true
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if getCount(i)+getCount(j) > k {
				continue
			}
			if i >= 1 {
				vis[i][j] = vis[i-1][j] || vis[i][j]
			}
			if j >= 1 {
				vis[i][j] = vis[i][j-1] || vis[i][j]
			}
			if vis[i][j] {
				result += 1
			}
		}
	}
	return result
}

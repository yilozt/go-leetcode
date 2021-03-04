package main

func spiralOrder(matrix [][]int) []int {
	var result []int
	if len(matrix) == 0 {
		return result
	}
	left, right, top, bottom := 0, len(matrix[0])-1, 0, len(matrix)-1
	for left <= right && top <= bottom {
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		for i := top + 1; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		if top != bottom {
			for i := right - 1; i >= left; i-- {
				result = append(result, matrix[bottom][i])
			}
		}
		if left != right {
			for i := bottom - 1; i > top; i-- {
				result = append(result, matrix[i][left])
			}
		}
		left += 1
		right -= 1
		top += 1
		bottom -= 1
	}
	return result
}

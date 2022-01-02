package main

// func findNumberIn2DArray(matrix [][]int, target int) bool {
// 	// 逐行二分，代价较大（不过写起来很简单XD）
// 	for _, row := range matrix {
// 		left, right := 0, len(row)-1
// 		for left <= right {
// 			mid := left + (right-left)/2
// 			if row[mid] == target {
// 				return true
// 			} else if row[mid] > target {
// 				right = mid - 1
// 			} else {
// 				left = mid + 1
// 			}
// 		}
// 	}
// 	return false
// }

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	rows, columns := len(matrix), len(matrix[0])
	row, column := 0, columns-1
	for row < rows && column >= 0 {
		if matrix[row][column] == target {
			return true
		} else if matrix[row][column] > target {
			column--
		} else {
			row++
		}
	}
	return false
}

package main

func findNumberIn2DArray(matrix [][]int, target int) bool {
	// 逐行二分，代价较大（不过写起来很简单XD）
	for _, row := range matrix {
		left, right := 0, len(row)-1
		for left <= right {
			mid := left + (right-left)/2
			if row[mid] == target {
				return true
			} else if row[mid] > target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return false
}

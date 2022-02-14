package main

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	var result []int
	bottom, right := len(matrix)-1, len(matrix[0])-1
	printCircle := func(start int) {
		for i := start; i <= right-start; i++ {
			result = append(result, matrix[start][i])
		}
		for i := start + 1; i <= bottom-start; i++ {
			result = append(result, matrix[i][right-start])
		}
		if start != bottom-start { // 上下的row不能相同
			for i := right - start - 1; i >= start; i-- {
				result = append(result, matrix[bottom-start][i])
			}
		}
		if start != right-start { // 左右的col不能相同
			for i := bottom - start - 1; i >= start+1; i-- {
				result = append(result, matrix[i][start])
			}
		}
	}
	var start int
	for bottom >= start*2 && right >= start*2 {
		printCircle(start)
		start++
	}
	return result
}

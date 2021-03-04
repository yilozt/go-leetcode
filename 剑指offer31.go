package main

func validateStackSequences(pushed []int, popped []int) bool {
	var stack []int
	var index int
	for _, v := range pushed {
		stack = append(stack, v)
		for len(stack) > 0 && stack[len(stack)-1] == popped[index] {
			index++
			stack = stack[:len(stack)-1]
		}
	}
	return index == len(popped)
}

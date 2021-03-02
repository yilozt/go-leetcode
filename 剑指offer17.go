package main

func printNumbers(n int) []int {
	var result []int
	for i := 1; i < quickPow(n); i++ {
		result = append(result, i)
	}
	return result
}

func quickPow(n int) int {
	result, tmp := 1, 10
	for n != 0 {
		if n&1 == 1 {
			result *= tmp
		}
		n >>= 1
		tmp *= tmp
	}
	return result
}

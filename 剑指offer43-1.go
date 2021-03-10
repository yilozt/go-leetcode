package main

func countDigitOne(n int) int {
	high := n / 10
	cur := n % 10
	low := 0
	digit := 1
	result := 0
	for !(high == 0 && cur == 0) {
		switch cur {
		case 0:
			result += high * digit
		case 1:
			result += high*digit + low + 1
		default:
			result += (high + 1) * digit
		}
		low += cur * digit
		cur = high % 10
		high /= 10
		digit *= 10
	}
	return result
}

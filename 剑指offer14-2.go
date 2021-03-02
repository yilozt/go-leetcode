package main

func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}
	result := 1
	const p = 1e9 + 7
	for n > 4 {
		result = result * 3 % p
		n -= 3
	}
	return result * n % p
}

package main

func myPow(x float64, n int) float64 {
	if n < 0 {
		x, n = 1/x, -n
	}
	// 使用快速幂
	var result float64 = 1
	tmp := x
	for n != 0 {
		if n&1 == 1 {
			result *= tmp
		}
		n >>= 1
		tmp *= tmp
	}
	return result
}

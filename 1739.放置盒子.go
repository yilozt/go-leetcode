package main

func minimumBoxes(n int) int {
	currentNum := 1
	height := 0
	i := 0
	for {
		if i*(i+1)*(i+2)/6 <= n {
			height = i
			currentNum = i * (i + 1) * (i + 2) / 6
		} else {
			break
		}
	}
	bottom := height * (height + 1) / 2
	sub := 0
	currentNum = n - currentNum
	for currentNum > 0 {
		bottom++
		sub++
		currentNum -= sub
	}
	return bottom
}

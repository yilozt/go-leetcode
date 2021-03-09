package main

func majorityElement(nums []int) int {
	var (
		vote   int
		result int
	)
	for _, v := range nums {
		if vote == 0 {
			result = v
		}
		if v == result {
			vote += 1
		} else {
			vote -= 1
		}
	}
	return result
}

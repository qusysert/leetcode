package main

func isPowerOfThree(n int) bool {
	num := float64(n)
	if num == 1 {
		return true
	}
	for num >= 3 {
		if num == 3 {
			return true
		}
		num = num / 3
	}
	return false
}

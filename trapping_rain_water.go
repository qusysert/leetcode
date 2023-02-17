package main

func trap(height []int) int {
	maxL := make(map[int]int)
	maxR := make(map[int]int)
	for i, v := range height {
		if i > 0 {
			for j := i - 1; j >= 0; j-- {
				if height[j] >= maxL[i] {
					maxL[i] = height[j]
					break
				}
			}
		} else {
			maxL[i] = v
		}
	}
	print(maxR)
	return 2
}

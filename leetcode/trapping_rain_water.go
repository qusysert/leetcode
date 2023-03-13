package main

func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}
	trapped := 0
	var maxL []int
	var maxR []int

	maxL = append(maxL, height[0])

	maxR = append(maxR, height[len(height)-1])
	for i := 1; i < len(height); i++ {
		if height[i] > maxL[i-1] {
			maxL = append(maxL, height[i])
		} else {
			maxL = append(maxL, maxL[i-1])
		}

		if height[len(height)-1-i] > maxR[i-1] {
			maxR = append(maxR, height[len(height)-1-i])
		} else {
			maxR = append(maxR, maxR[i-1])
		}
	}

	maxR = flip(maxR)

	for i := 0; i < len(height); i++ {
		blocks := min(maxL[i], maxR[i]) - height[i]
		if blocks >= 0 {
			trapped += blocks
		}
	}
	return trapped
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func flip(arr []int) []int {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
	return arr
}

package main

import "fmt"

func main() {
	fmt.Println(longestOnes([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3))
}

func longestOnes(nums []int, k int) int {
	begin := 0
	zeroCnt := 0
	maxLen := 0

	for end := range len(nums) {
		if nums[end] == 0 {
			zeroCnt++
		}

		for zeroCnt > k {
			if nums[begin] == 0 {
				zeroCnt--
			}
			begin++
		}

		maxLen = max(maxLen, end-begin+1)
	}

	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

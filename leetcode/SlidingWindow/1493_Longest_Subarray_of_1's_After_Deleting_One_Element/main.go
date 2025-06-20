package main

func longestSubarray(nums []int) int {
	begin := 0
	zeroCnt := 0
	maxLen := 0

	for end := range len(nums) {
		if nums[end] == 0 {
			zeroCnt++
		}

		for zeroCnt > 1 {
			if nums[begin] == 0 {
				zeroCnt--
			}
			begin++
		}

		maxLen = max(maxLen, end-begin+1)
	}

	return maxLen - 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

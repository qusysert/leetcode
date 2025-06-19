package main

import "math"

func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	cnt := 0
	left := 0
	right := len(nums) - 1

	for left <= right {
		if math.Abs(float64(nums[left])) > math.Abs(float64(nums[right])) {
			res[len(res)-1-cnt] = nums[left] * nums[left]
			left++
		} else {
			res[len(res)-1-cnt] = nums[right] * nums[right]
			right--
		}
		cnt++
	}
	return res
}

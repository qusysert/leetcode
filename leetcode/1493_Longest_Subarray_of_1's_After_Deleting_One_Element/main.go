// https://leetcode.com/problems/longest-subarray-of-1s-after-deleting-one-element/
package main

import "fmt"

func main() {
	fmt.Println(longestSubarray([]int{0, 1, 1, 1, 0, 1, 1, 0, 1}))
}

func longestSubarray(nums []int) int {
	l := 0
	r := 0
	max := 0
	prevZero := 0
	zeroCnt := 0
	if nums[0] == 0 {
		zeroCnt = 1
		prevZero = 0
	}
	for r < len(nums)-1 {
		r++
		if nums[r] == 0 {
			zeroCnt++
			if zeroCnt <= 1 {
				max = getMax(max, r-l+1)
				prevZero = r
			} else {
				l = prevZero + 1
				zeroCnt--
				prevZero = r
			}
		} else {
			max = getMax(max, r-l+1)
		}
	}
	if max == 0 {
		return 0
	}
	return max - 1
}

func getMax(a, b int) int {
	if a < b {
		return b
	}
	return a
}

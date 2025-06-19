package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	m := make(map[[3]int]struct{})
	var target int
	for i := range len(nums) {
		target = -nums[i]
		left := i + 1
		right := len(nums) - 1
		for left < right {
			cur := nums[left] + nums[right]

			switch {
			case cur == target:
				m[[3]int{-target, nums[left], nums[right]}] = struct{}{}
				left++
				right--
			case cur > target:
				right--
			case cur < target:
				left++
			}
		}
	}
	for k := range m {
		res = append(res, k[:])
	}
	return res
}

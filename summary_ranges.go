package main

import (
	"strconv"
)

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	var res []string
	cur := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		compare := nums[i-1] + 1
		if nums[i] == compare {
			cur = append(cur, nums[i])
		} else {
			if len(cur) > 1 {
				res = append(res, strconv.Itoa(cur[0])+"->"+strconv.Itoa(cur[len(cur)-1]))
			} else {
				res = append(res, strconv.Itoa(cur[0]))
			}

			cur = []int{nums[i]}
		}
	}
	if len(cur) > 1 {
		res = append(res, strconv.Itoa(cur[0])+"->"+strconv.Itoa(cur[len(cur)-1]))
	} else {
		res = append(res, strconv.Itoa(cur[0]))
	}
	return res
}

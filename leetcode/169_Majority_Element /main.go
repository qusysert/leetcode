package main

func majorityElement(nums []int) int {
	occurs := make(map[int]int)
	border := len(nums) / 2

	for _, v := range nums {
		occurs[v]++
		if occurs[v] > border {
			return v
		}
	}
	return 0
}

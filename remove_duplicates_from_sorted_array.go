package main

func RemoveDuplicates(nums []int) int {
	unique := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			unique++
		}
	}
	return unique
}

// https://leetcode.com/problems/remove-duplicates-from-sorted-array/
package main

func removeDuplicatesSorted(nums []int) int {
	unique := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			unique++
		}
	}
	return unique
}

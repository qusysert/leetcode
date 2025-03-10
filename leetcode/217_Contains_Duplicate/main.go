package main

func containsDuplicate(nums []int) bool {
	m := make(map[int]struct{}, len(nums))

	for i := range nums {
		if _, ok := m[nums[i]]; ok {
			return true
		}
		m[nums[i]] = struct{}{}
	}
	return false
}

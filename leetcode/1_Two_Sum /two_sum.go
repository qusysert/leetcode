// https://leetcode.com/problems/two-sum/
package __Two_Sum_

// for NOT sorted
func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, num := range nums {
		toAdd := target - num
		if idx, found := numMap[toAdd]; found {
			return []int{idx, i}
		}
		numMap[num] = i
	}
	return []int{}
}

// for sorted

//func twoSum(nums []int, target int) []int {
//	l := 0
//	r := len(nums) - 1
//	for l < r {
//		cur := nums[l] + nums[r]
//		if cur == target {
//			return []int{l, r}
//		}
//		if cur < target {
//			l++
//		}
//		if cur > target {
//			r--
//		}
//	}
//	return []int{}
//}

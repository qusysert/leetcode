// https://leetcode.com/problems/two-sum/
package __Two_Sum_

// for NOT sorted
func twoSum(nums []int, target int) []int {
	// val - position
	met := make(map[int]int)
	for i, v := range nums {
		if met[target-v] != 0 {
			return []int{met[target-v] - 1, i}
		}
		met[v] = i + 1
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

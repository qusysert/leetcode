// https://leetcode.com/problems/contains-duplicate-ii/
package main

import "math"

func main() {

}

func containsNearbyDuplicate(nums []int, k int) bool {
	indexes := make(map[int]int)
	for i, v := range nums {
		if idx, ok := indexes[v]; ok && int(math.Abs(float64(i-idx))) <= k {
			return true
		}
		indexes[v] = i
	}
	return false
}

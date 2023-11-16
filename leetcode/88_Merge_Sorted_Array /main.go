// https://leetcode.com/problems/merge-sorted-array
package main

import "fmt"

func main() {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	res := make([]int, m+n)
	cnt := m + n - 1
	m--
	n--

	for cnt >= 0 {
		if n < 0 {
			res[cnt] = nums1[m]
			cnt--
			m--
			continue
		}
		if m < 0 {
			res[cnt] = nums2[n]
			cnt--
			n--
			continue
		}
		if nums1[m] > nums2[n] {
			res[cnt] = nums1[m]
			m--
			cnt--

		} else if nums2[n] > nums1[m] {
			res[cnt] = nums2[n]
			n--
			cnt--

		} else {
			res[cnt] = nums1[m]
			cnt--
			res[cnt] = nums2[n]
			cnt--
			m--
			n--
		}
	}

	copy(nums1, res)
	fmt.Print(nums1)
}

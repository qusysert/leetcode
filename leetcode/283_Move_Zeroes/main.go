// https://leetcode.com/problems/move-zeroes/
package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Print(nums)
}

func moveZeroes(nums []int) {
	res := make([]int, len(nums))
	cnt := 0

	for _, num := range nums {
		if num != 0 {
			res[cnt] = num
			cnt++
		}
	}
	copy(nums, res)

}

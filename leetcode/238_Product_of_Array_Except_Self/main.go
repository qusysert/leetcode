package main

import "fmt"

func main() {
	fmt.Print(productExceptSelf([]int{-1, 1, 0, -3, 3}))
}

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))

	var prefix = 1

	for i := range nums {
		if i != 0 {
			prefix *= nums[i-1]
		}
		res[i] = prefix
	}

	var postfix = 1

	for i := len(nums) - 1; i >= 0; i-- {
		if i != len(nums)-1 {
			postfix *= nums[i+1]
		}
		res[i] *= postfix
	}

	return res
}

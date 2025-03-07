package main

import "fmt"

func main() {
	fmt.Println(topKFrequent([]int{4, 1, -1, 2, -1, 2, 3}, 2))
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int, len(nums))
	for i := range nums {
		m[nums[i]]++
	}
	freq := make([][]int, len(nums)+1)

	for num, amount := range m {
		if freq[amount] == nil {
			freq[amount] = []int{num}
		} else {
			freq[amount] = append(freq[amount], num)
		}
	}

	var res []int

	for i := len(freq) - 1; i >= 0; i-- {
		if k < 1 {
			break
		}
		if freq[i] != nil {
			res = append(res, freq[i]...)
			k -= len(freq[i])
		}
	}

	return res
}

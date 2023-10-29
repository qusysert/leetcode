package main

func main() {
	_ = longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1})
}

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	numsMap := make(map[int]bool)

	longest := 1

	for _, v := range nums {
		numsMap[v] = true
	}

	for _, v := range nums {
		cur := 1
		if !numsMap[v-1] {
			for numsMap[v+1] {
				cur++
				v++
			}
			if cur > longest {
				longest = cur
			}
		}
	}
	return longest
}

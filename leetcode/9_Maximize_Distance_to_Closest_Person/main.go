// https://leetcode.com/problems/maximize-distance-to-closest-person/description/
package main

import (
	"fmt"
)

func main() {
	fmt.Print(maxDistToClosest([]int{0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0}))
}

// 1, 0, 0, 0, 1, 0, 1
func maxDistToClosest(seats []int) int {

	maxDist := 0
	cur := 0

	// between 2 people
	for _, v := range seats {
		if v == 1 {
			cur = 0
		} else {
			cur++
			maxDist = getMax(maxDist, (cur+1)/2)

		}
	}

	// sitting left end
	cur = 0
	for _, v := range seats {
		if v == 1 {
			break
		} else {
			cur++
			maxDist = getMax(maxDist, cur)
		}
	}

	// sitting right end
	cur = 0
	for i := len(seats) - 1; i >= 0; i-- {
		if seats[i] == 1 {
			break
		} else {
			cur++
			maxDist = getMax(maxDist, cur)
		}
	}

	return maxDist
}

func getMax(a, b int) int {
	if a < b {
		return b
	}
	return a
}

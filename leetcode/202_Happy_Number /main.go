// https://leetcode.com/problems/happy-number
package main

import (
	"math"
	"strconv"
)

func main() {
	_ = isHappy(19)
}

func isHappy(n int) bool {
	calculated := make(map[int]bool)
	buf := n

	for {
		cur := calc(buf)
		if cur == 1 {
			return true
		}
		if _, ok := calculated[cur]; ok {
			return false
		}
		calculated[cur] = true
		buf = cur
	}
}

func calc(num int) int {
	var res int
	str := strconv.Itoa(num)
	for _, v := range str {
		digit, _ := strconv.Atoi(string(v))
		res += int(math.Pow(float64(digit), float64(2)))
	}
	return res
}

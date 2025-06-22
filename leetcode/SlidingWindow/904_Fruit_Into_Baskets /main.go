package main

import "fmt"

func main() {
	fmt.Println(totalFruit([]int{0, 1, 2, 2}))
}

func totalFruit(fruits []int) int {
	m := make(map[int]int)
	var begin int
	res := 0

	for end := range len(fruits) {
		m[fruits[end]]++

		for len(m) > 2 {
			if m[fruits[begin]] > 1 {
				m[fruits[begin]]--
			} else {
				delete(m, fruits[begin])
			}
			begin++
		}
		res = max(res, end-begin+1)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

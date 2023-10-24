// https://leetcode.com/problems/find-all-anagrams-in-a-string/
package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
}

func findAnagrams(s string, p string) []int {
	if len(p) > len(s) {
		return []int{}
	}
	var res []int
	windowMap := make(map[string]int)
	pMap := make(map[string]int)
	for i, _ := range p {
		windowMap[string(s[i])]++
		pMap[string(p[i])]++
	}
	if reflect.DeepEqual(windowMap, pMap) {
		res = append(res, 0)
	}
	l := 0
	for r := len(p); r < len(s); r++ {
		windowMap[string(s[r])]++
		windowMap[string(s[l])]--
		if windowMap[string(s[l])] == 0 {
			delete(windowMap, string(s[l]))
		}
		l++
		if reflect.DeepEqual(windowMap, pMap) {
			res = append(res, l)
		}
	}
	return res
}

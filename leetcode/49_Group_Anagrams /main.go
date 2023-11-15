// https://leetcode.com/problems/group-anagrams
package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, str := range strs {
		sortedStr := sortString(str)
		m[sortedStr] = append(m[sortedStr], str)
	}

	var res [][]string
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

func sortString(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool { return r[i] < r[j] })
	return string(r)
}

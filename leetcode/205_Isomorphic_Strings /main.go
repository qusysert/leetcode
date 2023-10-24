// https://leetcode.com/problems/isomorphic-strings/
package main

func isIsomorphic(s string, t string) bool {

	sMap := map[byte]int{}
	tMap := map[byte]int{}

	for i := range s {
		if sMap[s[i]] != tMap[t[i]] {
			return false
		} else {
			sMap[s[i]] = i + 1
			tMap[t[i]] = i + 1
		}
	}

	return true

}

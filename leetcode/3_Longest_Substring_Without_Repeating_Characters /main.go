// https://leetcode.com/problems/longest-substring-without-repeating-characters/
package main

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	chars := make(map[rune]int)
	maxLen := 0
	start := 1
	for i, c := range s {
		i++
		if chars[c] >= start {
			start = chars[c] + 1
		}
		chars[c] = i
		if i-start+1 > maxLen {
			maxLen = i - start + 1
		}
	}
	return maxLen
}

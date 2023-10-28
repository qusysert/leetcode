// https://leetcode.com/problems/valid-palindrome/
package _25_Valid_Palindrome_

import (
	"strings"
)

func IsPalindrome(s string) bool {
	if strings.ToLower(reverseString1([]byte(clearString(s)))) == strings.ToLower(clearString(s)) {
		return true
	}
	return false
}

func clearString(s string) string {
	var result strings.Builder
	for i := 0; i < len(s); i++ {
		b := s[i]
		if ('a' <= b && b <= 'z') ||
			('A' <= b && b <= 'Z') ||
			('0' <= b && b <= '9') {
			result.WriteByte(b)
		}
	}
	return result.String()
}

func reverseString1(s []byte) string {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
	return string(s)
}

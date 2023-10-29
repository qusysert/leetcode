// https://leetcode.com/problems/word-pattern/description/
package main

import (
	"strings"
)

func main() {
	_ = wordPattern("abba", "dog cat cat dog")
}

func wordPattern(pattern string, s string) bool {
	wordMapPattern := make(map[string]int)
	wordMapString := make(map[string]int)
	stringArr := strings.Split(s, " ")

	if len(stringArr) != len(pattern) {
		return false
	}

	sCnt := 0
	pCnt := 0

	for i := 0; i < len(pattern); i++ {
		if _, ok := wordMapPattern[string(pattern[i])]; !ok {
			wordMapPattern[string(pattern[i])] = pCnt
			pCnt++
		}

		if _, ok := wordMapString[stringArr[i]]; !ok {
			wordMapString[stringArr[i]] = sCnt
			sCnt++
		}

	}
	if pCnt != sCnt {
		return false
	}

	for i := 0; i < len(pattern); i++ {
		patternWord := string(pattern[i])
		strWord := stringArr[i]
		strWordId := wordMapString[strWord]
		patternWordId := wordMapPattern[patternWord]

		if strWordId != patternWordId {
			return false
		}
	}
	return true
}

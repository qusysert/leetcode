package main

import "fmt"

func main() {

	fmt.Print(lengthOfLastWord("a"))
}

func lengthOfLastWord(s string) int {
	var endOfLastWord = 0
	for i := len(s) - 1; i >= 0; i-- {
		if rune(s[i]) != 32 {
			endOfLastWord = i
			break
		}
	}

	cnt := 0
	for i := endOfLastWord; i >= 0; i-- {

		if rune(s[i]) == 32 {
			return cnt
		}
		cnt++
	}
	return cnt
}

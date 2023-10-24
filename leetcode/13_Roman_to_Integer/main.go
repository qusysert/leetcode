// https://leetcode.com/problems/roman-to-integer/
package main

import "fmt"

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	conv := map[uint8]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	res := 0
	prev := 0

	for i := len(s) - 1; i >= 0; i-- {
		if conv[s[i]] >= prev {
			res += conv[s[i]]
		} else {
			res -= conv[s[i]]
		}
		prev = conv[s[i]]
	}
	return res
}

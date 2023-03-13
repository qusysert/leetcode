package main

func IsAnagram(s string, t string) int {
	chars := make(map[rune]int)
	for _, c := range s {
		chars[c]++
	}
	for _, c := range t {
		if chars[c] <= 0 {
			return 0
		}
		chars[c]--
	}
	for _, element := range chars {
		if element > 0 {
			return 0
		}
	}
	return 1
}

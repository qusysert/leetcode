// https://leetcode.com/problems/letter-combinations-of-a-phone-number/description/
package main

func main() {

}

var dial = map[rune][]string{
	'2': []string{"a", "b", "c"},
	'3': []string{"d", "e", "f"},
	'4': []string{"g", "h", "i"},
	'5': []string{"j", "k", "l"},
	'6': []string{"m", "n", "o"},
	'7': []string{"p", "q", "r", "s"},
	'8': []string{"t", "u", "v"},
	'9': []string{"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	if len(digits) == 1 {
		return dial[rune(digits[0])]
	}

	res := dial[rune(digits[0])]
	for i := 1; i < len(digits); i++ {
		res = getComb(res, dial[rune(digits[i])])
	}
	return res
}

func getComb(first []string, second []string) []string {
	var result []string
	for i := 0; i < len(first); i++ {
		for j := 0; j < len(second); j++ {
			result = append(result, first[i]+second[j])
		}
	}
	return result
}

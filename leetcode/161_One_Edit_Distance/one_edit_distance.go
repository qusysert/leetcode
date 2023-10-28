// https://wentao-shao.gitbook.io/leetcode/string/161.one-edit-distance
package _61_One_Edit_Distance

func oneEditDistance(s string, t string) bool {
	if len(t) > len(s) {
		return oneEditDistance(t, s)
	}
	if len(s)-len(t) != 1 {
		return false
	}
	for i := 0; i < len(t); i++ {
		if s[i] != t[i] {
			if len(s) == len(t) {
				return s[i+1:] == t[i+1:]
			} else {
				return s[i+1:] == t[i:]
			}
		}
	}
	return true
}

// https://leetcode.com/problems/valid-parentheses/description/
package _0_Valid_Parentheses_

func isValid(s string) bool {
	if len(s) == 0 || len(s) == 1 || len(s)%2 != 0 {
		return false
	}
	var stack Stack
	pair := make(map[string]string)
	pair[string(')')] = string('(')
	pair[string('}')] = string('{')
	pair[string(']')] = string('[')
	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			stack.Push(string(c))
		}
		if c == ')' || c == ']' || c == '}' {
			last, _ := stack.Pop()
			if last != pair[string(c)] {
				return false
			}
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

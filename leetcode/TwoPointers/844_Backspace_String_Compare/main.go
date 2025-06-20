package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(backspaceCompare("ab#c", "ad#c"))
}

func backspaceCompare(s string, t string) bool {
	p1 := len(s) - 1
	p2 := len(t) - 1

	var res1, res2 []string
	var cnt1, cnt2 int

	for p1 >= 0 {
		if s[p1] == '#' {
			cnt1++
			p1--
			continue
		}

		if cnt1 > 0 {
			cnt1--
			p1--
		} else {
			res1 = append(res1, string(s[p1]))
			p1--
		}
	}

	for p2 >= 0 {
		if t[p2] == '#' {
			cnt2++
			p2--
			continue
		}

		if cnt2 > 0 {
			cnt2--
			p2--
		} else {
			res2 = append(res2, string(t[p2]))
			p2--
		}
	}

	fmt.Println(res1, res2)
	return reflect.DeepEqual(res1, res2)
}

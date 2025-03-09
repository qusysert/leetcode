package main

import (
	"fmt"
	"strconv"
)

const delimiter = "#"

func main() {
	str := encode([]string{"we", "say", ":", "yes", "!@#$%^&*()"})
	fmt.Println(decode(str))
}

func encode(arr []string) string {
	var res string
	for i := range arr {
		res += strconv.Itoa(len(arr[i])) + delimiter + arr[i]
	}
	return res
}

func decode(arr string) []string {
	var res []string
	var curLen int

	pos := 0

	for {
		curLenStr := ""

		if pos+1 > len(arr) {
			break
		}

		for {
			if string(arr[pos]) == delimiter {
				break
			}
			curLenStr += string(arr[pos])
			pos++
		}
		curLen, _ = strconv.Atoi(curLenStr)
		res = append(res, arr[pos+1:pos+1+curLen])
		pos = pos + curLen + 1
	}
	return res
}

// 20#so3lee
// 0123456

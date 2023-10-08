// https://leetcode.com/problems/string-compression/description/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	strArray := []string{"a", "b", "c"}
	byteArray := make([]byte, len(strArray))

	for i, str := range strArray {
		byteArray[i] = str[0]
	}
	fmt.Print(compress(byteArray))
}

func compress(chars []byte) int {
	compressed := make([]byte, 0, len(chars))
	var prevC byte
	curCnt := 0
	if len(chars) == 1 {
		return 1
	}
	for i, c := range chars {
		if i == 0 {
			prevC = c
			curCnt++
			continue
		}
		if i == len(chars)-1 {
			if c == prevC {
				compressed = append(compressed, prevC)
				compressed = append(compressed, []byte(strconv.Itoa(curCnt+1))...)
				continue
			} else {
				compressed = append(compressed, prevC)
				if curCnt > 1 {
					compressed = append(compressed, []byte(strconv.Itoa(curCnt+2))...)
				}
				compressed = append(compressed, c)
			}
			break
		}
		if c == prevC {
			curCnt++
			continue
		}

		compressed = append(compressed, prevC)
		if curCnt > 1 {
			compressed = append(compressed, []byte(strconv.Itoa(curCnt))...)
		}
		curCnt = 1
		prevC = c

	}
	copy(chars, compressed)
	return len(compressed)
}

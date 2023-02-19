package main

import "fmt"

//
//import (
//	"bufio"
//	"fmt"
//	"os"
//	"strconv"
//	"strings"
//)
//
//func main() {
//	r := bufio.NewReader(os.Stdin)
//
//	nKeysStr := read(*r)
//	nKeys, _ := strconv.Atoi(nKeysStr)
//
//	symbols := stringToArray(read(*r))
//	rowCoordinated := stringToArray(read(*r))
//	lengthStr := read(*r)
//	length, _ := strconv.Atoi(lengthStr)
//	text := stringToArray(read(*r))
//
//	fmt.Println(countPower(nKeys, symbols, rowCoordinated, length, text))
//
//}
//
//func countPower(nKeys int, symbols []int, rowCoordinated []int, length int, text []int) int {
//	power := 0
//	keyboard := map[int]int{}
//	for i := 0; i < len(rowCoordinated); i++ {
//		row := rowCoordinated[i]
//		symbol := symbols[i]
//		keyboard[symbol] = row
//	}
//	var prevRow int
//	var curRow int
//	for i := 0; i < length; i++ {
//		curSymbolId := text[i]
//		curRow = keyboard[curSymbolId]
//		if prevRow != 0 && curRow != prevRow {
//			power++ //+= int(math.Abs(float64(curRow - prevRow)))
//		}
//		prevRow = curRow
//	}
//	return power
//}
//
//func read(r bufio.Reader) string {
//	s, _ := r.ReadString('\n')
//	return strings.Replace(s, "\n", "", -1)
//}
//
//func stringToArray(str string) []int {
//	strArr := strings.Split(str, " ")
//	arr := make([]int, len(strArr))
//	for i, v := range strArr {
//		arr[i], _ = strconv.Atoi(v)
//	}
//	return arr
//}

func main() {
	fmt.Println(trap([]int{4, 2, 0, 3, 2, 5}))
}

// https://leetcode.com/problems/pascals-triangle/
package main

import "fmt"

func main() {
	fmt.Println(generate(5))
}

func generate(numRows int) [][]int {
	var triangle [][]int
	triangle = append(triangle, []int{1})
	var lastRow []int

	for i := 0; i < numRows-1; i++ {
		lastRow = triangle[len(triangle)-1]
		var newRow []int
		newRow = append(newRow, 1)
		if len(lastRow) >= 1 {
			for j := 0; j < len(lastRow)-1; j++ {
				newRow = append(newRow, lastRow[j]+lastRow[j+1])
			}
		}
		newRow = append(newRow, 1)
		triangle = append(triangle, newRow)
	}
	return triangle
}

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var ways []int

	reader := bufio.NewReader(os.Stdin)
	dimensionsInput, _ := reader.ReadString('\n')
	dimensions := strings.Fields(strings.TrimSpace(dimensionsInput))
	y, _ := strconv.Atoi(dimensions[0])
	x, _ := strconv.Atoi(dimensions[1])

	field := make([][]int, y)
	for i := 0; i < y; i++ {
		line, _ := reader.ReadString('\n')
		numbers := strings.Fields(strings.TrimSpace(line))
		field[i] = make([]int, x)
		for j := 0; j < x; j++ {
			field[i][j], _ = strconv.Atoi(numbers[j])
		}
	}

	calc(field[0][0], [2]int{0, 0}, field, &ways)

	fmt.Print(getMin(ways))
}

func calc(total int, coords [2]int, field [][]int, ways *[]int) {
	x := coords[0]
	y := coords[1]
	lenY := len(field[0])
	lenX := len(field)

	if coords == [2]int{lenX - 1, lenY - 1} {
		*ways = append(*ways, total)
		return
	}

	if lenX > 1 && x+1 <= lenX-1 && total+field[x+1][y] <= getMin(*ways) {
		calc(total+field[x+1][y], [2]int{x + 1, y}, field, ways)
	}
	if lenY > 1 && y+1 <= lenY-1 && total+field[x][y+1] <= getMin(*ways) {
		calc(total+field[x][y+1], [2]int{x, y + 1}, field, ways)
	}
	return
}

func getMin(arr []int) int {
	shortest := math.MaxInt64
	for _, num := range arr {
		if num < shortest {
			shortest = num
		}
	}
	return shortest
}

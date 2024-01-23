package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reflect([][]int{{1, 1}, {-1, 1}}))
}

func reflect(points [][]int) bool {
	var minX = math.MaxInt
	var maxX = math.MinInt
	var pointMap = make(map[[2]int]struct{})

	for i := 0; i < len(points); i++ {
		// Putting coords into map for further checking of points
		pointMap[[2]int{points[i][0], points[i][1]}] = struct{}{}

		// Getting min and max Xs
		if points[i][0] < minX {
			minX = points[i][0]
		}
		if points[i][0] > maxX {
			maxX = points[i][0]
		}
	}

	var lineX = (maxX + minX) / 2
	fmt.Printf("Reflect line X coord: %d\n", lineX)

	for point, _ := range pointMap {
		// Calculating a reflection for every point depending on it's position relative to lineX
		var reflectedPoint [2]int
		if point[0] > lineX {
			reflectedPoint = [2]int{lineX - (point[0] - lineX), point[1]}
		}
		if point[0] < lineX {
			reflectedPoint = [2]int{lineX + (point[0] + lineX), point[1]}
		}
		// If point is on the line - there is no reflection for it
		if point[0] == lineX {
			return false
		}

		fmt.Printf("Original: %v\n", point)
		fmt.Printf("Reflected: %v\n", reflectedPoint)

		// Checking if there is a reflected point for all points
		if _, ok := pointMap[reflectedPoint]; !ok {
			return false
		}
	}
	return true

}

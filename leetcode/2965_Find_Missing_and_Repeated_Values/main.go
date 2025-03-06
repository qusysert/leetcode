package main

func main() {

}

func findMissingAndRepeatedValues(grid [][]int) []int {
	m := make(map[int]int, len(grid)*len(grid[0]))

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			m[grid[i][j]]++
		}
	}

	var missing, repeated int

	for i := 1; i <= len(grid)*len(grid[0]); i++ {
		switch m[i] {
		case 0:
			missing = i
		case 2:
			repeated = i
		}
		if missing != 0 && repeated != 0 {
			return []int{missing, repeated}
		}
	}
	return []int{missing, repeated}
}

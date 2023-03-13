package main

func bubble(arr []int) []int {
	if len(arr) == 0 || len(arr) == 1 {
		return arr
	}
	for {
		isSorted := true
		for i := 1; i < len(arr); i++ {
			if arr[i] < arr[i-1] {
				arr[i], arr[i-1] = arr[i-1], arr[i]
				isSorted = false
			}
		}
		if isSorted {
			return arr
		}
	}
}

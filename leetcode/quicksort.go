package main

func quickSort(arr []int) []int {
	return sort(arr, 0, len(arr)-1)
}

func sort(arr []int, start, end int) []int {
	if end <= start {
		return arr
	}
	pivot := partition(arr, start, end)
	sort(arr, 0, pivot-1)
	sort(arr, pivot+1, end)
	return arr
}

func partition(arr []int, start, end int) int {
	pivot := arr[end]
	i := start - 1
	for j := start; j < end; j++ {
		if arr[j] < pivot {
			i++
			arr[j], arr[i] = arr[i], arr[j]
		}
	}
	i++
	arr[end], arr[i] = arr[i], arr[end]
	return i
}

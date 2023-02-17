package main

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	left := mergeSort(arr[:len(arr)/2])
	right := mergeSort(arr[len(arr)/2:])
	return merge(left, right)
}

func merge(leftArr, rightArr []int) []int {
	l := 0
	r := 0
	var sorted []int
	for l < len(leftArr) && r < len(rightArr) {
		if leftArr[l] < rightArr[r] {
			sorted = append(sorted, leftArr[l])
			l++
		} else {
			sorted = append(sorted, rightArr[r])
			r++
		}
	}

	for ; l < len(leftArr); l++ {
		sorted = append(sorted, leftArr[l])
	}

	for ; r < len(rightArr); r++ {
		sorted = append(sorted, rightArr[r])
	}

	return sorted
}

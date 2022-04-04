package binarysearch

import "fmt"

//BinarySearch binary search algoritm
func BinarySearch(item int, arr []int) (int, bool) {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2
		result := arr[mid]
		if result == item {
			return mid, true
		}
		if result > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return 0, false
}

func RunBinarySearch() {

	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	i, b := BinarySearch(10, array)
	fmt.Println(i, " ", b)
}

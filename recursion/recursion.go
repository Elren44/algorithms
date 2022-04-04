package recursion

import (
	"fmt"
)

func Factorial(x int) int {
	if x == 1 {
		return 1
	} else {
		return x * Factorial(x-1)
	}
}

func sum(arr []int) int {
	var result int
	if len(arr) == 0 {
		return 0
	} else if len(arr) == 1 {
		return arr[0]
	} else {
		result = arr[0] + sum(arr[1:])
	}

	return result
}

func RunSum() {
	arr := []int{1, 2, 3, 4, 5, 5}
	sum := sum(arr)
	fmt.Println(sum)
}

func RunBinarySearchRecursion() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	i, b := binarySearchRecursion(2, array)
	fmt.Println(i, " ", b)
}

func binarySearchRecursion(item int, arr []int) (result int, found bool) {
	mid := len(arr) / 2

	switch {
	case len(arr) == 0:
		result = -1 // not found
	case arr[mid] > item:
		result, found = binarySearchRecursion(item, arr[:mid])
	case arr[mid] < item:
		result, found = binarySearchRecursion(item, arr[mid+1:])
		if result >= 0 { // if anything but the -1 "not found" result
			result += mid + 1
		}
	default: // a[mid] == search
		result = mid // found
		found = true
	}
	return result, found
}

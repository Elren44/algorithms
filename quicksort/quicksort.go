package quicksort

import (
	"fmt"
)

func quickSort(arr []int) {

	if len(arr) < 2 {
		return
	}

	p := partition(arr)

	quickSort(arr[:p])
	quickSort(arr[p+1:])

}

func partition(arr []int) int {
	if len(arr) < 2 {
		return 0
	}
	low := 0
	high := len(arr) - 1
	pivot := arr[len(arr)/2]
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[j], arr[low] = arr[low], arr[j]
			low++
		}
	}
	arr[low], arr[high] = arr[high], arr[low]
	return low
}

func RunQuickSort() {
	arr := []int{56, 13, 1, 4, 2, 77, 12, 24, 33, 21, 11, 3, 5, 15, 99, 87}
	fmt.Println(arr)
	arr = MyTestSort(arr)
	//quickSort(arr)
	//sort.Ints(arr)
	fmt.Println(arr)

}

func MyTestSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	index := len(arr) / 2
	p := arr[index]
	less := []int{}
	gr := []int{}
	for i := 0; i < len(arr); i++ {
		if arr[i] <= p && i != index {
			less = append(less, arr[i])
		}
		if arr[i] > p && i != index {
			gr = append(gr, arr[i])
		}
	}
	quickSort(less)
	arr = []int{}
	arr = append(less, p)
	quickSort(gr)
	arr = append(arr, gr...)
	return arr
}

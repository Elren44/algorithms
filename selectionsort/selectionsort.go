package selectionsort

import "fmt"

//SelectionSort golang example
func goSelectionSort(arr []int) {
	var size = len(arr)
	for i := 0; i < size; i++ {
		min := i
		for j := i; j < size; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
}

func RunSelectionSort() {
	arr := []int{2, 104, 25, 22, 14, 3, 1, 65, 17}
	fmt.Println(arr)
	//goSelectionSort(arr)
	arr = selectionSort(arr)
	fmt.Println(arr)
}

//Grokking Algorithms example

func findSmallest(arr []int) int {
	min := arr[0]
	min_index := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
			min_index = i
		}
	}
	return min_index
}

func selectionSort(arr []int) []int {
	newArr := []int{}
	size := len(arr)
	for i := 0; i < size; i++ {
		min := findSmallest(arr)
		newArr = append(newArr, arr[min])
		arr = append(arr[:min], arr[min+1:]...)
	}

	return newArr
}

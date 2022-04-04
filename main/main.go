package main

import (
	"fmt"
	"time"

	"github.com/Elren44/algoritms/quicksort"
)

func main() {
	t := time.Now()
	quicksort.RunQuickSort()
	durations := time.Since(t)
	fmt.Println(durations.Nanoseconds())
	// arr := []int{56, 13, 1, 4, 2, 77, 12, 24, 33, 21, 11, 3, 5, 15, 99, 87}
	// fmt.Println(arr)

	// //sort.Ints(arr)
	// fmt.Println(arr)
}

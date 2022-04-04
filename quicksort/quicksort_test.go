package quicksort

import "testing"

func BenchmarkRunQuickSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RunQuickSort()
	}
}

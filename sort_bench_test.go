package sort

import (
	gsort "sort"
	"testing"
)

func BenchmarkStdLibSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := []int{5, 2, 6, 3, 1, 4, 3, 5, 3, 1, 2, 7, 8, 3}
		gsort.Ints(s)
	}
}

func BenchmarkHeapSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := []int{5, 2, 6, 3, 1, 4, 3, 5, 3, 1, 2, 7, 8, 3}
		s = HeapSort(s)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := []int{5, 2, 6, 3, 1, 4, 3, 5, 3, 1, 2, 7, 8, 3}
		QuickSort(s)
	}
}

package sort

import (
	gsort "sort"
	"testing"
)

func BenchmarkStdLibSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := createNewSlice(numElement)
		gsort.Ints(s)
	}
}

func BenchmarkHeapSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := createNewSlice(numElement)
		s = HeapSort(s)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := createNewSlice(numElement)
		QuickSort(s)
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := createNewSlice(numElement)
		BubbleSort(s)
	}
}

func BenchmarkSorterInterface(b *testing.B) {
	for n := 0; n < b.N; n++ {
		c := []cat{
			cat{"addie"},
			cat{"moritz"},
			cat{"roger"},
			cat{"duwey"},
			cat{"lorrie"},
			cat{"jen"},
			cat{"annie"},
		}

		Quick(cats(c))
	}
}

func BenchmarkStdLibSorterInterface(b *testing.B) {
	for n := 0; n < b.N; n++ {
		c := []cat{
			cat{"addie"},
			cat{"moritz"},
			cat{"roger"},
			cat{"duwey"},
			cat{"lorrie"},
			cat{"jen"},
			cat{"annie"},
		}

		gsort.Sort(cats(c))
	}
}

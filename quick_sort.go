//Wikipedia definition
//Quicksort is a divide and conquer algorithm which relies on a partition operation:
//to partition an array an element called a pivot is selected.
//All elements smaller than the pivot are moved before it and all greater elements are moved after it.
//This can be done efficiently in linear time and in-place.
//The lesser and greater sublists are then recursively sorted.
//This yields average time complexity of O(n log n), with low overhead, and thus this is a popular algorithm.
//Efficient implementations of quicksort (with in-place partitioning) are typically unstable sorts and somewhat complex, but are among the fastest sorting algorithms in practice.
//Together with its modest O(log n) space usage, quicksort is one of the most popular sorting algorithms and is available in many standard programming libraries.

package sort

//quicksort algorithm based on current high and low part of the array
func quickSort(inp []int, lo, hi int) {
	//return if lo is already same with hi
	if lo >= hi {
		return
	}

	//pivot start with the highest index
	pivot := hi
	//loop from the start of the slice until the highest index minus the current pivot
	for i := lo; i < pivot; i++ {
		//compare the pivot value with the value of index i
		for inp[i] >= inp[pivot] && i < pivot {
			//if it's greater than the current pivot, swap the value
			inp[i], inp[pivot], inp[pivot-1] = inp[pivot-1], inp[i], inp[pivot]
			//decrement pivot by 1
			pivot--
		}
	}
	quickSort(inp, lo, pivot-1)
	quickSort(inp, pivot+1, hi)
}

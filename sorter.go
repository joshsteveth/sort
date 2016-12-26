//interface sorter enable all type of data to enable this sort package
package sort

type Sorter interface {
	//return the length of the slice
	Len() int
	//swap element i and j of the slice
	Swap(i, j int)
	//return whether i is lesser than j
	Less(i, j int) bool
}

//Since quick sort until now is the fastest here
//utilize the interface to do quick sort
func Quick(s Sorter) {
	quick(s, 0, s.Len()-1)
}

func quick(s Sorter, lo, hi int) {
	//return if low is already greater or equal the high index
	if lo >= hi {
		return
	}

	//pivot start with the highest index
	pivot := hi
	//loop from the start of the slice until the highest index minus the current pivot
	for i := lo; i < pivot; i++ {
		//compare the pivot value with the value of index i
		for s.Less(pivot, i) && i < pivot {
			//if it's greater than the current pivot, swap the value
			s.Swap(i, pivot-1)
			s.Swap(pivot, pivot-1)
			//decrement pivot by 1
			pivot--
		}
	}
	quick(s, lo, pivot-1)
	quick(s, pivot+1, hi)
}

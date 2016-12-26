##Collection of sorting algorith
###This is only for exercise, not fast or anything special (std library is 100x faster :p)

###Current features:
- Heap sort
- Quick sort

Benchmark Result:
```
//Go standard Library
BenchmarkStdLibSort-4   	 3000000	       514 ns/op
//Heap Sort
BenchmarkHeapSort-4     	   30000	     47618 ns/op
//Quick Sort (it's actually faster)
BenchmarkQuickSort-4    	10000000	       158 ns/op
```

##Heap Sort
```
foo := []int{6, 5, 3, 1, 8, 7, 2, 4, 4, 5}
foo = sort.HeapSort(foo)
//now foo should be: []int{1, 2, 3, 4, 4, 5, 5, 6, 7, 8}
```

##Quick Sort
```
foo := []int{6, 5, 3, 1, 8, 7, 2, 4, 4, 5}
sort.QuickSort(foo)
//now foo should be: []int{1, 2, 3, 4, 4, 5, 5, 6, 7, 8}
```
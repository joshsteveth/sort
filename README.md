##Collection of sorting algorith
###This is only for exercise, not very fast or anything special 

###Current features:
- Heap sort
- Quick sort

Benchmark Result: //create slice with 200 random elem + sort
```
//go std lib
BenchmarkStdLibSort-4   	   50000	     32386 ns/op
//heap sort (certainly poorly written by me lol)
BenchmarkHeapSort-4     	     200	   6130642 ns/op
//quick sort
BenchmarkQuickSort-4    	  100000	     21035 ns/op
//bubble sort
BenchmarkBubbleSort-4   	   20000	     93965 ns/op
```

##Heap Sort
```
foo := []int{6, 5, 3, 1, 8, 7, 2, 4, 4, 5}
foo = sort.HeapSort(foo)
```

##Quick Sort
```
foo := []int{6, 5, 3, 1, 8, 7, 2, 4, 4, 5}
sort.QuickSort(foo)
```

##Bubble Sort
```
foo := []int{6, 5, 3, 1, 8, 7, 2, 4, 4, 5}
sort.BubbleSort(foo)
```

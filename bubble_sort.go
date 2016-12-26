//Bubble sort is a simple sorting algorithm.
//The algorithm starts at the beginning of the data set.
//It compares the first two elements, and if the first is greater than the second, it swaps them.
//It continues doing this for each pair of adjacent elements to the end of the data set.
//It then starts again with the first two elements, repeating until no swaps have occurred on the last pass.
//This algorithm's average time and worst-case performance is O(n2), so it is rarely used to sort large, unordered data sets.
//Bubble sort can be used to sort a small number of items (where its asymptotic inefficiency is not a high penalty).

package sort

func BubbleSort(inp []int) {
	var sorted bool

	for !sorted {
		sorted = bubbleSort(inp)
	}
}

//return whether input is already sorted or not
func bubbleSort(inp []int) bool {
	alreadySorted := true

	for i := 0; i < len(inp)-1; i++ {
		if inp[i] > inp[i+1] {
			//swap the value of both index
			//also set the sorted to be true
			inp[i], inp[i+1] = inp[i+1], inp[i]
			alreadySorted = false
		}
	}
	return alreadySorted
}

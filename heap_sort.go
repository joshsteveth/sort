//Wikipedia definition:
//Heapsort is a much more efficient version of selection sort. 
//It also works by determining the largest (or smallest) element of the list, 
//placing that at the end (or beginning) of the list, then continuing with the rest of the list, 
//but accomplishes this task efficiently by using a data structure called a heap, a special type of binary tree. 
//Once the data list has been made into a heap, the root node is guaranteed to be the largest (or smallest) element. 
//When it is removed and placed at the end of the list, 
//the heap is rearranged so the largest element remaining moves to the root. 
//Using the heap, finding the next largest element takes O(log n) time, 
//instead of O(n) for a linear scan as in simple selection sort. 
//This allows Heapsort to run in O(n log n) time, and this is also the worst case complexity.

package sort

//heap has a value, parent, and left/right child
type heap struct {
	value      int
	parent     *heap
	leftChild  *heap
	rightChild *heap
}

//utilize the heap sort
func HeapSort(inp []int) []int {
	newArr := []int{}
	length := len(inp)

	for len(newArr) != length {
		//create new heap and rearrange it with heap sort rule
		h := newHeap(inp)

		//turn back heap h into []int
		lvl := h.getMaxLevel()
		arr := h.reArray(lvl)

		//get the biggest heap and the new array as well
		//append the biggest value into new Array
		var biggest int
		inp, biggest = reArrangeHeap(arr)

		newArr = append(newArr, biggest)
	}

	return newArr
}

func newHeap(inp []int) *heap {
	//set first value as the highest heap hierarchy at first
	h := &heap{
		value: inp[0],
	}

	//create map to determine which level of heap we currently in
	currentHeap := h
	currentLevel := 1

	heapMap := map[int][]*heap{
		currentLevel: []*heap{currentHeap},
	}

	for i := 1; i < len(inp); i++ {
		newHeap := currentHeap.addNewVal(inp[i])
		newHeap.swap()

		//register the new Heap to the map
		heapMap[currentLevel+1] = append(heapMap[currentLevel+1], newHeap)

		//change current heap into new one
		currentHeap, currentLevel = getCurrentHeap(heapMap, currentHeap, currentLevel)
	}

	return h
}

func (h *heap) addNewVal(val int) *heap {
	newHeap := heap{
		value:  val,
		parent: h,
	}

	//check whether the left child is already occupied or not
	//if yes then occupy the right child
	if h.leftChild == nil {
		h.leftChild = &newHeap
	} else {
		h.rightChild = &newHeap
	}

	return &newHeap
}

//return the new current heap and level
func getCurrentHeap(mp map[int][]*heap, currentHeap *heap, currentLevel int) (*heap, int) {
	//check the current heap right child first
	//if it's still empty then simply return the current heap
	if currentHeap.rightChild == nil {
		return currentHeap, currentLevel
	}

	//use the heap map
	//check whether all left and right child of heaps on this level is already occupied or not
	for key, val := range mp[currentLevel] {
		//if either left child or right child of this map is not fully occupied
		//return the heap and current level
		if val.leftChild == nil || val.rightChild == nil {
			return mp[currentLevel][key], currentLevel
		}
	}

	//if all of them are already occupied then return the first left child of current level
	//increment current level by 1
	return mp[currentLevel][0].leftChild, currentLevel + 1
}

//swap the value with parent if it is lesser
//parent should always have lesser value than its child
func (h *heap) swap() {
	//return if this heap does not any parent
	if h.parent == nil {
		return
	}

	if h.value <= h.parent.value {
		h.value, h.parent.value = h.parent.value, h.value
		//also continue the swap for the parent as well
		h.parent.swap()
	}
}

//get max level of the heap
func (h *heap) getMaxLevel() int {
	totalLvl := 1
	curHeap := h
	for curHeap.leftChild != nil {
		totalLvl++
		curHeap = curHeap.leftChild
	}
	return totalLvl
}

//change back the heap into map[int][]*heap
func (h *heap) reverseMap(maxLvl int) map[int][]*heap {
	//create new heapMap with the max level
	heapMap := map[int][]*heap{}
	for i := 1; i <= maxLvl; i++ {
		heapMap[i] = []*heap{}
	}

	//update the map for parent
	heapMap[1] = append(heapMap[1], h)

	//only update for max lvl -1
	for i := 1; i < maxLvl; i++ {
		//update the map for all heaps on this level
		for _, hp := range heapMap[i] {
			//update only if left/right heap is not nil
			if hp.leftChild != nil {
				heapMap[i+1] = append(heapMap[i+1], hp.leftChild)
			}
			if hp.rightChild != nil {
				heapMap[i+1] = append(heapMap[i+1], hp.rightChild)
			}
		}
	}

	return heapMap
}

//change back the heap into []int
func (h *heap) reArray(maxLvl int) []int {
	newArr := []int{}

	//use the reverseMap func first
	rMap := h.reverseMap(maxLvl)

	for i := 1; i <= maxLvl; i++ {
		for _, v := range rMap[i] {
			newArr = append(newArr, v.value)
		}
	}

	return newArr
}

//switch the first and last element of an array
//return the last element and create new slice without it
func reArrangeHeap(inp []int) ([]int, int) {
	biggestNum := inp[0]
	//switch the 1st and last value of inp
	inp[0], inp[len(inp)-1] = inp[len(inp)-1], inp[0]

	//create new slice without its last element
	newArr := []int{}
	for i := 0; i < len(inp)-1; i++ {
		newArr = append(newArr, inp[i])
	}

	return newArr, biggestNum
}

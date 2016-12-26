package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testingArray = []int{6, 5, 3, 1, 8, 7, 2, 4}

func TestNewHeap(t *testing.T) {
	h := newHeap(testingArray)
	assert.Equal(t, 1, h.value)
	assert.Equal(t, 3, h.leftChild.value)
	assert.Equal(t, 2, h.rightChild.value)
	assert.Equal(t, 4, h.leftChild.leftChild.value)
	assert.Equal(t, 8, h.leftChild.rightChild.value)
	assert.Equal(t, 7, h.rightChild.leftChild.value)
	assert.Equal(t, 5, h.rightChild.rightChild.value)
	assert.Equal(t, 6, h.leftChild.leftChild.leftChild.value)
}

func TestHeapGetMaxLevel(t *testing.T) {
	h := newHeap(testingArray)
	lvl := h.getMaxLevel()
	assert.Equal(t, 4, lvl)
}

func TestHeapReverseMapAndArray(t *testing.T) {
	h := newHeap(testingArray)
	lvl := h.getMaxLevel()
	heapMap := h.reverseMap(lvl)
	assert.Equal(t, 1, len(heapMap[1]))
	assert.Equal(t, 2, len(heapMap[2]))
	assert.Equal(t, 4, len(heapMap[3]))
	assert.Equal(t, 1, len(heapMap[4]))
}

func TestHeapReArray(t *testing.T) {
	h := newHeap(testingArray)
	lvl := h.getMaxLevel()
	newArr := h.reArray(lvl)
	expected := []int{1, 3, 2, 4, 8, 7, 5, 6}
	assert.Equal(t, expected, newArr)
}

func TestRearrangeHeap(t *testing.T) {
	inp := []int{8, 6, 7, 4, 5, 3, 2, 1}
	newArr, biggest := reArrangeHeap(inp)
	expected := []int{1, 6, 7, 4, 5, 3, 2}
	assert.Equal(t, expected, newArr)
	assert.Equal(t, 8, biggest)
}

func TestHeapSort(t *testing.T) {
	foo := []int{6, 5, 3, 1, 8, 7, 2, 4, 4, 5}
	foo = HeapSort(foo)
	expected := []int{1, 2, 3, 4, 4, 5, 5, 6, 7, 8}
	assert.Equal(t, expected, foo)
}

package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	var testQSArray = []int{3, 7, 8, 5, 2, 1, 9, 5, 4}

	QuickSort(testQSArray)
	expected := []int{1, 2, 3, 4, 5, 5, 7, 8, 9}
	assert.Equal(t, expected, testQSArray)
}

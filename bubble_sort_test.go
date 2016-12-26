package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	inp := []int{5, 1, 4, 2, 8}
	BubbleSort(inp)
	expected := []int{1, 2, 4, 5, 8}
	assert.Equal(t, expected, inp)
}

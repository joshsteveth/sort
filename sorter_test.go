package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type cat struct {
	name string
}

type cats []cat

func (c cats) Len() int           { return len(c) }
func (c cats) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c cats) Less(i, j int) bool { return c[i].name < c[j].name }

func TestQuickSorterInterface(t *testing.T) {
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
	assert.Equal(t, "addie", c[0].name)
	assert.Equal(t, "annie", c[1].name)
	assert.Equal(t, "duwey", c[2].name)
	assert.Equal(t, "jen", c[3].name)
	assert.Equal(t, "lorrie", c[4].name)
	assert.Equal(t, "moritz", c[5].name)
	assert.Equal(t, "roger", c[6].name)
}

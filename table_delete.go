package yti

import (
	"slices"
)

func (t *Table[I]) Delete(matches func(item I) bool) {
	for rowNum, item := range t.items {
		if matches(item) {
			t.removeItemAtRow(rowNum)
		}
	}
}

func (t *Table[I]) removeItemAtRow(rowNum int) {
	t.items = slices.Delete(t.items, rowNum, rowNum+1)
	t.dirty = true
	t.recalculateIndex(rowNum)
}

package yti

func (t *Table[I]) Insert(item I) error {
	t.items = append(t.items, item)
	t.dirty = true
	// recalculate all the indices
	return t.indexItem(item, len(t.items)-1)
}

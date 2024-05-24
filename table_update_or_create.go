package yti

// Update items based on `matches` function or create a new one
func (t *Table[I]) UpdateOrCreate(item I, matches func(item I) bool) error {
	var foundItems bool

	for rowNum, existingItem := range t.items {
		if matches(existingItem) {
			// found item, let's update it
			t.replaceItem(item, rowNum)
			foundItems = true
		}
	}

	if !foundItems {
		return t.Insert(item)
	}

	return nil
}

func (t *Table[I]) replaceItem(item I, rowNum int) {
	t.items[rowNum] = item
	t.reindexItem(item, rowNum)
}

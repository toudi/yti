package yti

func (t *Table[I]) GetOrCreate(lookup func(item I) bool, newItem I) (I, error) {
	for _, existingItem := range t.items {
		if lookup(existingItem) {
			return existingItem, nil
		}
	}

	return newItem, t.Insert(newItem)
}

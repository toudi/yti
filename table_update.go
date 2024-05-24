package yti

func (t *Table[I]) Update(updated func(item I) bool) {
	for row, item := range t.items {
		if updated(item) {
			t.reindexItem(item, row)
		}
	}
}

package yti

// calls `visitor` functior for each visited item. if `visitor` returns true, the function will stop iterating
func (t *Table[I]) ForEach(visitor func(item I) bool) {
	for _, item := range t.items {
		if visitor(item) {
			break
		}
	}
}

// return items that `matches` user-defined criteria
func (t *Table[I]) Fetch(matches func(item I) bool) []I {
	var results []I

	t.ForEach(func(item I) bool {
		if matches(item) {
			results = append(results, item)
		}
		return false
	})

	return results
}

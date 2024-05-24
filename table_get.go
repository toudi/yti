package yti

func (t *Table[I]) Get(matches func(item I) bool) (I, error) {
	var rowsMatched int
	var rowNum int

	for row, item := range t.items {
		if matches(item) {
			rowNum = row
			rowsMatched += 1
		}
	}

	if rowsMatched == 0 {
		return t.zeroValue, ErrItemDoesNotExist
	}

	if rowsMatched > 1 {
		return t.zeroValue, ErrRowsNotUnique
	}

	return t.items[rowNum], nil
}

func (t *Table[I]) First(matches func(item I) bool) (I, error) {
	for _, item := range t.items {
		if matches(item) {
			return item, nil
		}
	}

	return t.zeroValue, ErrItemDoesNotExist
}

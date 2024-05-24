package yti

// looks up values by index `indexName`. Returns the first one
// - if the index does not exist, it returns ErrUnknownIndex
// - if there's more than one row with the specified `value`, returns ErrRowsNotUnique
func (t *Table[I]) GetByIndex(indexName string, value interface{}) (I, error) {
	rows, err := t.rowsByIndexValue(indexName, value)
	if err != nil {
		return t.zeroValue, err
	}

	if len(rows) == 1 {
		return t.items[rows[0]], nil
	}
	if len(rows) > 1 {
		return t.zeroValue, ErrRowsNotUnique
	}
	return t.zeroValue, ErrItemDoesNotExist
}

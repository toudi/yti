package yti

func (t *Table[I]) GetOrCreateByIndexValue(
	indexName string,
	indexValue interface{},
	newItem I,
) (I, error) {
	rows, err := t.rowsByIndexValue(indexName, indexValue)
	if err != nil {
		return t.zeroValue, err
	}

	if len(rows) == 0 {
		return newItem, t.Insert(newItem)
	}

	if len(rows) > 1 {
		return t.zeroValue, ErrRowsNotUnique
	}

	return t.items[rows[0]], nil
}

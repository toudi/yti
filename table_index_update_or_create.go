package yti

func (t *Table[I]) UpdateOrCreateByIndexValue(
	indexName string,
	indexValue interface{},
	newItem I,
) error {
	rows, err := t.rowsByIndexValue(indexName, indexValue)
	// unknown index
	if err != nil {
		return err
	}

	// no items found - we can create new one
	if len(rows) == 0 {
		return t.Insert(newItem)
	}

	// we have to update rows one by one
	for _, row := range rows {
		t.replaceItem(newItem, row)
	}

	return nil
}

func (t *Table[I]) UpdateOrCreateByIndexValues(indexName string, newItems map[interface{}]I) error {
	for indexValue, newItem := range newItems {
		if err := t.UpdateOrCreateByIndexValue(indexName, indexValue, newItem); err != nil {
			return err
		}
	}

	return nil
}

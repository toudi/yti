package yti

func (t *Table[I]) DeleteByIndexValue(indexName string, indexValue interface{}) error {
	rows, err := t.rowsByIndexValue(indexName, indexValue)
	if err != nil {
		return err
	}

	for _, rowNum := range rows {
		t.removeItemAtRow(rowNum)
	}

	return nil
}

package yti

// this function was inspired by redis's SET NX - it creates an item if one does not exist.
// I think it only makes sense within the scope of the index, otherwise you'd have to
// iterate all the values to know that such item does not (yet) exist.
func (t *Table[I]) CreateNXByIndexValue(
	indexName string,
	indexValue interface{},
	newItem I,
) (bool, error) {
	rows, err := t.rowsByIndexValue(indexName, indexValue)
	if err != nil {
		// unknown index
		return false, err
	}

	if len(rows) == 0 {
		// item for the specified index value did not exist yet
		return true, t.Insert(newItem)
	}

	// item already in the table.
	return false, nil
}

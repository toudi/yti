package yti

import "errors"

var ErrCheckLimitCountExceeded = errors.New("checkCountLimit exceeded")

func (t *Table[I]) EnsureIndexDoesNotContain(
	indexName string,
	nextValueFunc func() interface{},
	checkCountLimit int,
) (interface{}, error) {
	if _, err := t.getIndex(indexName); err != nil {
		return nil, err
	}

	for i := 0; i < checkCountLimit; i += 1 {
		nextValue := nextValueFunc()
		rows, err := t.rowsByIndexValue(indexName, nextValue)
		if err != nil {
			return nil, err
		}
		if len(rows) == 0 {
			return nextValue, nil
		}
	}

	return nil, ErrCheckLimitCountExceeded
}

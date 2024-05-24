package yti

import "github.com/samber/lo"

func (t *Table[I]) FetchByIndexValue(indexName string, indexValue interface{}) ([]I, error) {
	rows, err := t.rowsByIndexValue(indexName, indexValue)
	if err != nil {
		return nil, err
	}

	return lo.Map(rows, func(rowNum int, _ int) I { return t.items[rowNum] }), nil
}

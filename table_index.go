package yti

import (
	"fmt"

	"github.com/samber/lo"
)

type Index map[string][]int

// because the user might want to index items by any arbitrary value,
// we're converting it to string. Otherwise, the following scenario would
// fail:
// a struct that has a field which is an uint8 which is then searched against int with the same value
func indexValue(value interface{}) string {
	return fmt.Sprintf("%v", value)
}

func (t *Table[I]) getIndex(name string) (Index, error) {
	var zero Index
	if index, exists := t.indices[name]; exists {
		return index, nil
	}
	// we don't have any rows for that index but it's stil possible that this is a fresh table
	// so let's check if this is a valid index
	if _, exists := t.options.Indices[name]; exists {
		return zero, nil
	}
	return zero, ErrUnknownIndex
}

func (t *Table[I]) rowsByIndexValue(indexName string, value interface{}) ([]int, error) {
	index, err := t.getIndex(indexName)
	if err != nil {
		return nil, err
	}

	return index[indexValue(value)], nil
}

func (t *Table[I]) indexItems() error {
	for row, item := range t.items {
		if err := t.indexItem(item, row); err != nil {
			return err
		}
	}

	return nil
}

func (t *Table[I]) indexItem(item I, row int) error {
	if t.options == nil {
		return nil
	}
	if t.indices == nil {
		t.indices = make(map[string]Index)
	}

	for indexName, indexer := range t.options.Indices {
		if _, exists := t.indices[indexName]; !exists {
			t.indices[indexName] = make(Index)
		}
		t.indices[indexName][indexValue(indexer(item))] = append(
			t.indices[indexName][indexValue(indexer(item))],
			row,
		)
	}

	return nil
}

func (t *Table[I]) reindexItem(item I, rowNo int) {
	for indexName, indexer := range t.options.Indices {
		for indexValue, indexedRows := range t.indices[indexName] {
			if lo.Contains(indexedRows, rowNo) {
				t.indices[indexName][indexValue] = lo.DropWhile(
					indexedRows,
					func(item int) bool { return item == rowNo },
				)
			}
		}
		newIndexValue := indexValue(indexer(item))
		t.indices[indexName][newIndexValue] = append(t.indices[indexName][newIndexValue], rowNo)
	}
}

func (t *Table[I]) recalculateIndex(removedRowNum int) {
	// the idea here is very simple. this function is called when we
	// remove item from collection. this means that for each row that
	// is *greater* than the removed row, we have to decrease it's value
	// in the index. for the rows that are less the removed one we don't
	// have to do anything. By using this technique we don't have to
	// reindex items which could be potentially expensive.
	for indexName, indexValues := range t.indices {
		for lookupValue, rows := range indexValues {
			var translatedRows []int
			for _, row := range rows {
				if row < removedRowNum {
					translatedRows = append(translatedRows, row)
				} else if row == removedRowNum {
					continue
				} else {
					translatedRows = append(translatedRows, row-1)
				}
			}
			t.indices[indexName][lookupValue] = translatedRows
		}
	}
}

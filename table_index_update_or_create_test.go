package yti_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/toudi/yti"
)

func TestUpdateOrCreateByIndex(t *testing.T) {
	t.Run("update or create - create path", func(t *testing.T) {
		t.Parallel()
		table := InMemoryTable(t, nil)
		item := TestStruct{Id: 1, Name: "test"}
		assert.NoError(t, table.UpdateOrCreateByIndexValue(IndexId, 1, item))
	})
	t.Run("update or create - update path", func(t *testing.T) {
		t.Parallel()
		var item = TestStruct{Id: 1, Name: "test"}
		table := InMemoryPopulatedTable(t, []TestStruct{item}, nil)

		assert.NoError(
			t,
			table.UpdateOrCreateByIndexValue(IndexId, 1, item),
		)
	})

	t.Run("update or create - invalid index", func(t *testing.T) {
		t.Parallel()
		var item = TestStruct{Id: 1, Name: "test"}
		table := InMemoryPopulatedTable(t, []TestStruct{item}, nil)

		assert.Equal(
			t,
			yti.ErrUnknownIndex,
			table.UpdateOrCreateByIndexValue("invalid", 1, item),
		)
	})
}

func TestUpdateOrCreateByIndexValues(t *testing.T) {
	t.Run("update or create by index values - create path", func(t *testing.T) {
		t.Parallel()
		table := InMemoryTable(t, nil)
		item := TestStruct{Id: 1, Name: "test"}
		assert.NoError(
			t,
			table.UpdateOrCreateByIndexValues(IndexId, map[interface{}]TestStruct{1: item}),
		)
	})
	t.Run("update or create by index values - update path", func(t *testing.T) {
		t.Parallel()
		table := InMemoryTable(t, nil)
		item := TestStruct{Id: 1, Name: "test"}
		_ = table.Insert(item)
		assert.NoError(
			t,
			table.UpdateOrCreateByIndexValues(IndexId, map[interface{}]TestStruct{1: item}),
		)
	})
	t.Run("update or create by index values - invalid index", func(t *testing.T) {
		t.Parallel()
		table := InMemoryTable(t, nil)
		item := TestStruct{Id: 1, Name: "test"}
		_ = table.Insert(item)
		assert.Equal(
			t,
			yti.ErrUnknownIndex,
			table.UpdateOrCreateByIndexValues("invalid", map[interface{}]TestStruct{1: item}),
		)
	})
}

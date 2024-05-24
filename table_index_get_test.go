package yti_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/toudi/yti"
)

func TestGetByIndex(t *testing.T) {
	t.Run("get by index - item does not exist", func(t *testing.T) {
		t.Parallel()

		table := InMemoryTable(t, nil)
		var zeroStruct = TestStruct{}
		result, err := table.GetByIndex(IndexId, 1)
		assert.Equal(t, yti.ErrItemDoesNotExist, err)
		assert.Equal(t, zeroStruct, result)
	})

	t.Run("get by index - item exists", func(t *testing.T) {
		t.Parallel()

		var expectedItem = TestStruct{Id: 1, Name: "test"}
		table := InMemoryTable(t, nil)
		assert.NoError(t, table.Insert(expectedItem))
		var rowNum uint8 = 1
		result, err := table.GetByIndex(IndexId, rowNum)
		assert.NoError(t, err)
		assert.Equal(t, expectedItem, result)
	})

	t.Run("get by index - unknown item", func(t *testing.T) {
		t.Parallel()

		table := InMemoryPopulatedTable(t, []TestStruct{{Id: 1, Name: "test"}}, nil)
		_, err := table.GetByIndex(IndexId, 0)
		assert.Equal(t, yti.ErrItemDoesNotExist, err)
	})

	t.Run("get by index - unknown index", func(t *testing.T) {
		t.Parallel()

		table := InMemoryPopulatedTable(t, []TestStruct{{Id: 1, Name: "test"}}, nil)
		_, err := table.GetByIndex("invalid", 1)
		assert.Equal(t, yti.ErrUnknownIndex, err)
	})

	t.Run("get by index - non unique rows", func(t *testing.T) {
		t.Parallel()

		table := InMemoryPopulatedTable(
			t,
			[]TestStruct{{Id: 1, Name: "test"}, {Id: 1, Name: "test"}},
			nil,
		)
		_, err := table.GetByIndex(IndexId, 1)
		assert.Equal(t, yti.ErrRowsNotUnique, err)
	})

}

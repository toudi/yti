package yti_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/toudi/yti"
)

func TestDeleteByIndexValue(t *testing.T) {
	t.Run("delete by index value (empty table)", func(t *testing.T) {
		t.Parallel()

		table := InMemoryTable(t, nil)
		assert.NoError(
			t,
			table.DeleteByIndexValue(IndexId, 1),
		)
	})
	t.Run("delete by index value - happy path", func(t *testing.T) {
		t.Parallel()

		table := InMemoryTable(t, nil)
		assert.NoError(t, table.Insert(TestStruct{Id: 1}))
		assert.NoError(
			t,
			table.DeleteByIndexValue(IndexId, 1),
		)
	})

	t.Run("delete by index value - invalid index", func(t *testing.T) {
		t.Parallel()

		table := InMemoryTable(t, nil)
		assert.NoError(t, table.Insert(TestStruct{Id: 1}))
		assert.Equal(
			t,
			yti.ErrUnknownIndex,
			table.DeleteByIndexValue("invalid index", 1),
		)
	})

	t.Run(
		"make sure that looking up by index still works after delete",
		func(t *testing.T) {
			t.Parallel()

			table := InMemoryPopulatedTable(
				t,
				[]TestStruct{{Id: 1, Name: "test 1"}, {Id: 2, Name: "test 2"}},
				nil,
			)

			assert.NoError(t, table.DeleteByIndexValue(IndexId, 1))
			entry, err := table.GetByIndex(IndexId, 2)
			assert.NoError(t, err)
			assert.Equal(t, TestStruct{Id: 2, Name: "test 2"}, entry)
		},
	)
}

package yti_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/toudi/yti"
)

func TestGetOrCreateByInde(t *testing.T) {
	t.Run("get or create with index - item not present", func(t *testing.T) {
		t.Parallel()
		lookup := &Lookup{}
		expectedStruct := TestStruct{Id: 1, Name: "test"}
		table, err := yti.OpenFile[TestStruct]("", &yti.TableOptions[TestStruct]{
			Indices: TestStructIndices,
		})
		assert.NoError(t, err)
		assert.NotNil(t, table)
		item, err := table.GetOrCreateByIndexValue(IndexId, expectedStruct.Id, expectedStruct)
		assert.NoError(t, err)
		assert.Equal(t, expectedStruct, item)
		lookup.AssertNotCalled(t, "itemFound", expectedStruct)
	})

	t.Run("get or create with index - item present", func(t *testing.T) {
		t.Parallel()
		lookup := &Lookup{}
		table := InMemoryTable(t, nil)
		expectedStruct := TestStruct{Id: 1, Name: "test"}
		assert.NoError(t, table.Insert(expectedStruct))
		item, err := table.GetOrCreateByIndexValue(IndexId, expectedStruct.Id, expectedStruct)
		assert.NoError(t, err)
		assert.Equal(t, expectedStruct, item)
		lookup.AssertNotCalled(t, "itemFound", expectedStruct)
	})

	t.Run("get or create - unknown index", func(t *testing.T) {
		t.Parallel()

		table := InMemoryTable(t, nil)
		expectedStruct := TestStruct{Id: 1, Name: "test"}
		_, err := table.GetOrCreateByIndexValue("invalid", 1, expectedStruct)
		assert.Equal(t, yti.ErrUnknownIndex, err)
	})
	t.Run("get or create - non unique rows", func(t *testing.T) {
		t.Parallel()

		expectedStruct := TestStruct{Id: 1, Name: "test"}
		table := InMemoryPopulatedTable(t, []TestStruct{expectedStruct, expectedStruct}, nil)
		_, err := table.GetOrCreateByIndexValue(IndexId, 1, expectedStruct)
		assert.Equal(t, yti.ErrRowsNotUnique, err)
	})

}

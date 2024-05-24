package yti_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/toudi/yti"
)

func TestCreateNx(t *testing.T) {
	t.Run("createnx - create path", func(t *testing.T) {
		t.Parallel()
		table := InMemoryTable(t, nil)
		created, err := table.CreateNXByIndexValue(IndexId, 1, TestStruct{Id: 1, Name: "test"})
		assert.NoError(t, err)
		assert.True(t, created)
	})

	t.Run("createnx - invalid index", func(t *testing.T) {
		t.Parallel()
		table := InMemoryTable(t, nil)
		created, err := table.CreateNXByIndexValue("invalid", 1, TestStruct{Id: 1, Name: "test"})
		assert.Equal(t, yti.ErrUnknownIndex, err)
		assert.False(t, created)
	})

	t.Run("createnx - item already on the list", func(t *testing.T) {
		t.Parallel()
		var expected = TestStruct{Id: 1, Name: "test"}
		table := InMemoryPopulatedTable(t, []TestStruct{expected}, nil)
		created, err := table.CreateNXByIndexValue(IndexId, 1, expected)
		assert.NoError(t, err)
		assert.False(t, created)
	})
}

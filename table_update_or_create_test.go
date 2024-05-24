package yti_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateOrCreate(t *testing.T) {
	t.Run("update or create - create path", func(t *testing.T) {
		t.Parallel()
		table := InMemoryTable(t, nil)
		item := TestStruct{Id: 1, Name: "test"}
		assert.NoError(t, table.UpdateOrCreate(item, func(item TestStruct) bool { return false }))
	})
	t.Run("update or create - update path", func(t *testing.T) {
		t.Parallel()
		table := InMemoryTable(t, nil)
		item := TestStruct{Id: 1, Name: "test"}
		_ = table.Insert(item)
		assert.NoError(
			t,
			table.UpdateOrCreate(item, func(entry TestStruct) bool { return entry.Id == item.Id }),
		)
	})
}

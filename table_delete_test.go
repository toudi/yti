package yti_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDelete(t *testing.T) {
	t.Run("delete on empty table", func(t *testing.T) {
		t.Parallel()

		table := InMemoryTable(t, nil)
		table.Delete(func(item TestStruct) bool { return false })
		assert.Equal(
			t,
			0,
			table.Count(),
		)
	})

	t.Run("delete - happy path", func(t *testing.T) {
		t.Parallel()

		table := InMemoryTable(t, nil)
		assert.NoError(t, table.Insert(TestStruct{Id: 1}))
		table.Delete(func(item TestStruct) bool { return item.Id == 1 })
		assert.Equal(
			t,
			0,
			table.Count(),
		)
	})
	t.Run("delete - more than one item in the seed", func(t *testing.T) {
		t.Parallel()

		var expected = TestStruct{Id: 2, Name: "test"}
		table := InMemoryPopulatedTable(
			t,
			[]TestStruct{{Id: 1, Name: "test"}, expected, {Id: 3, Name: "test"}},
			nil,
		)
		table.Delete(func(item TestStruct) bool { return item.Id == expected.Id })
		assert.Equal(
			t,
			2,
			table.Count(),
		)
	})
}

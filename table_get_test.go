package yti_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/toudi/yti"
)

func TestTableGet(t *testing.T) {
	t.Run("get on empty table", func(t *testing.T) {
		t.Parallel()

		table := InMemoryTable(t, nil)
		_, err := table.Get(func(item TestStruct) bool { return false })
		assert.Equal(
			t,
			yti.ErrItemDoesNotExist,
			err,
		)
	})
	t.Run("get - non-unique items", func(t *testing.T) {
		t.Parallel()

		table := InMemoryPopulatedTable(
			t,
			[]TestStruct{{Id: 1, Name: "test"}, {Id: 2, Name: "test"}},
			nil,
		)
		_, err := table.Get(func(item TestStruct) bool { return item.Name == "test" })
		assert.Equal(
			t,
			yti.ErrRowsNotUnique,
			err,
		)
	})
	t.Run("get - happy path", func(t *testing.T) {
		t.Parallel()

		var expected = TestStruct{Id: 2, Name: "test"}

		table := InMemoryPopulatedTable(
			t,
			[]TestStruct{{Id: 1, Name: "test"}, expected},
			nil,
		)
		item, err := table.Get(func(item TestStruct) bool { return item.Id == expected.Id })
		assert.NoError(t, err)
		assert.Equal(
			t,
			expected,
			item,
		)
	})
}

func TestTableFirst(t *testing.T) {
	t.Run("first - empty table", func(t *testing.T) {
		t.Parallel()

		table := InMemoryTable(t, nil)
		_, err := table.First(func(item TestStruct) bool { return true })
		assert.Equal(t, yti.ErrItemDoesNotExist, err)
	})
	t.Run("first - happy path", func(t *testing.T) {
		t.Parallel()

		var expected = TestStruct{Id: 2, Name: "lookup"}
		table := InMemoryPopulatedTable(
			t,
			[]TestStruct{{Id: 1, Name: "test"}, expected, {Id: 3, Name: "lookup"}},
			nil,
		)
		item, err := table.First(func(item TestStruct) bool { return item.Name == expected.Name })
		assert.Equal(t, expected, item)
		assert.NoError(t, err)
	})
}

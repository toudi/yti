package yti_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForEach(t *testing.T) {
	t.Run("for each - function does not return false", func(t *testing.T) {
		t.Parallel()

		table := InMemoryTable(t, nil)
		_ = table.Insert(TestStruct{Id: 1, Name: "test 1"})
		_ = table.Insert(TestStruct{Id: 2, Name: "test 2"})
		_ = table.Insert(TestStruct{Id: 3, Name: "test 3"})

		lookup := &Lookup{}
		lookup.On("itemFound", TestStruct{Id: 1, Name: "test 1"}).Return(false)
		lookup.On("itemFound", TestStruct{Id: 2, Name: "test 2"}).Return(false)
		lookup.On("itemFound", TestStruct{Id: 3, Name: "test 3"}).Return(false)

		table.ForEach(lookup.itemFound)
		lookup.AssertNumberOfCalls(t, "itemFound", 3)
	})

	t.Run("for each - function stops after first item", func(t *testing.T) {
		t.Parallel()

		table := InMemoryTable(t, nil)
		_ = table.Insert(TestStruct{Id: 1, Name: "test 1"})
		_ = table.Insert(TestStruct{Id: 2, Name: "test 2"})
		_ = table.Insert(TestStruct{Id: 3, Name: "test 3"})

		lookup := &Lookup{}
		lookup.On("itemFound", TestStruct{Id: 1, Name: "test 1"}).Return(true)

		table.ForEach(lookup.itemFound)
		lookup.AssertNumberOfCalls(t, "itemFound", 1)
	})
}

func TestFetch(t *testing.T) {
	t.Run("fetch all items", func(t *testing.T) {
		t.Parallel()

		table := InMemoryTable(t, nil)
		_ = table.Insert(TestStruct{Id: 1, Name: "test 1"})
		_ = table.Insert(TestStruct{Id: 2, Name: "test 2"})

		entries := table.Fetch(func(item TestStruct) bool { return true })
		assert.Equal(t, []TestStruct{{Id: 1, Name: "test 1"}, {Id: 2, Name: "test 2"}}, entries)
	})

	t.Run("fetch selected items", func(t *testing.T) {
		t.Parallel()

		table := InMemoryTable(t, nil)
		_ = table.Insert(TestStruct{Id: 1, Name: "foo"})
		_ = table.Insert(TestStruct{Id: 2, Name: "bar"})

		entries := table.Fetch(func(item TestStruct) bool { return item.Name == "foo" })
		assert.Equal(t, []TestStruct{{Id: 1, Name: "foo"}}, entries)
	})
}

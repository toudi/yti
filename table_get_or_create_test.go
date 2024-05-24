package yti_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type Lookup struct {
	mock.Mock
}

func (l *Lookup) itemFound(item TestStruct) bool {
	args := l.Called(item)
	return args.Bool(0)
}

func TestGetOrCreate(t *testing.T) {
	t.Run("get or create - item not present", func(t *testing.T) {
		t.Parallel()
		lookup := &Lookup{}

		expectedStruct := TestStruct{Id: 1, Name: "test"}
		table := InMemoryTable(t, nil)
		item, err := table.GetOrCreate(lookup.itemFound, expectedStruct)
		assert.NoError(t, err)
		assert.Equal(t, expectedStruct, item)
		lookup.AssertNotCalled(t, "itemFound", expectedStruct)
	})

	t.Run("get or create - item not present", func(t *testing.T) {
		t.Parallel()
		lookup := &Lookup{}
		expectedStruct := TestStruct{Id: 1, Name: "test"}
		table := InMemoryPopulatedTable(t, []TestStruct{expectedStruct}, nil)
		lookup.On("itemFound", expectedStruct).Return(true)
		item, err := table.GetOrCreate(lookup.itemFound, expectedStruct)
		assert.NoError(t, err)
		assert.Equal(t, expectedStruct, item)
		lookup.AssertExpectations(t)
	})

}

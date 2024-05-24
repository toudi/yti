package yti_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/toudi/yti"
)

func TestFetchByIndexValue(t *testing.T) {
	t.Run("fetch by index value - happy path", func(t *testing.T) {
		t.Parallel()

		table := InMemoryPopulatedTable(
			t,
			[]TestStruct{{Id: 1, Name: "test"}, {Id: 2, Name: "lookup"}, {Id: 3, Name: "lookup"}},
			&yti.TableOptions[TestStruct]{
				Indices: map[string]yti.Indexer[TestStruct]{
					"name": func(item TestStruct) interface{} {
						return item.Name
					},
				},
			},
		)

		items, err := table.FetchByIndexValue("name", "lookup")
		assert.NoError(t, err)
		assert.Equal(t, []TestStruct{{Id: 2, Name: "lookup"}, {Id: 3, Name: "lookup"}}, items)
	})
	t.Run("fetch by index value - invalid index", func(t *testing.T) {
		t.Parallel()

		table := InMemoryPopulatedTable(
			t,
			[]TestStruct{{Id: 1, Name: "test"}, {Id: 2, Name: "lookup"}, {Id: 3, Name: "lookup"}},
			nil,
		)

		_, err := table.FetchByIndexValue("name", "lookup")
		assert.Equal(t, yti.ErrUnknownIndex, err)
	})
}

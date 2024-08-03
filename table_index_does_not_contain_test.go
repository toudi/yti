package yti_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/toudi/yti"
)

func TestEnsureIndexDoesNotContain(t *testing.T) {
	t.Run("empty table - should return value", func(t *testing.T) {
		t.Parallel()

		table := InMemoryTable(t, nil)

		id, err := table.EnsureIndexDoesNotContain("id", func() interface{} { return 1 }, 1)
		assert.NoError(
			t,
			err,
		)
		assert.Equal(t, 1, id)
	})

	t.Run(
		"populated table - should return error about forloop counter exceeding",
		func(t *testing.T) {
			t.Parallel()

			table := InMemoryPopulatedTable(
				t,
				[]TestStruct{{Id: 1}},
				nil,
			)

			id, err := table.EnsureIndexDoesNotContain("id", func() interface{} { return 1 }, 1)

			assert.Equal(t, nil, id)
			assert.Equal(t, yti.ErrCheckLimitCountExceeded, err)
		},
	)

	t.Run("empty table - no indices defined", func(t *testing.T) {
		t.Parallel()

		table := InMemoryTable(t, nil)

		id, err := table.EnsureIndexDoesNotContain(
			"non-existing",
			func() interface{} { return 1 },
			1,
		)
		assert.Equal(t, nil, id)
		assert.Equal(t, yti.ErrUnknownIndex, err)
	})
}

package yti_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/toudi/yti"
)

func TestUpdate(t *testing.T) {
	t.Run("make sure that reindex item was called", func(t *testing.T) {
		fetchBySomPropFalse := func(item *TestStruct) bool {
			return item.SomeProp == false
		}

		WithTempFile(func(file *os.File) {
			_, _ = file.WriteString("---\n")
			table, err := yti.OpenFile[*TestStruct](file.Name(), &yti.TableOptions[*TestStruct]{
				Indices: map[string]yti.Indexer[*TestStruct]{
					"by-prop": func(item *TestStruct) interface{} {
						return item.SomeProp
					},
				},
			})
			assert.NoError(t, err)
			assert.NoError(t, table.Insert(&TestStruct{Id: 1, Name: "test", SomeProp: true}))
			assert.Len(t, table.Fetch(fetchBySomPropFalse), 0)
			table.Update(func(item *TestStruct) bool {
				if item.SomeProp == true {
					item.SomeProp = false
					return true
				}
				return false
			})
			assert.Len(t, table.Fetch(fetchBySomPropFalse), 1)
		})
	})
}

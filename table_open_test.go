package yti_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/toudi/yti"
)

func TestOpenFile(t *testing.T) {
	t.Run("open nonexisting file should not yield errors by default", func(t *testing.T) {
		t.Parallel()
		_, err := yti.OpenFile[TestStruct]("non-existing-file", nil)
		assert.NoError(t, err)
	})

	t.Run("open nonexisting file with mustExist option should yield error", func(t *testing.T) {
		t.Parallel()
		_, err := yti.OpenFile[TestStruct]("non-existing-file", &yti.TableOptions[TestStruct]{
			MustExist: true,
		})
		assert.True(t, os.IsNotExist(err))
	})

	t.Run("open existing file (invalid yaml)", func(t *testing.T) {
		t.Parallel()
		WithTempFile(func(file *os.File) {
			_, err := yti.OpenFile[TestStruct](
				file.Name(),
				&yti.TableOptions[TestStruct]{MustExist: true},
			)
			assert.Equal(t, yti.ErrLoadingTable, err)
		})
	})

	t.Run("open existing file (empty yaml)", func(t *testing.T) {
		t.Parallel()
		WithTempFile(func(file *os.File) {
			_, _ = file.WriteString("---\n")
			file.Close()
			_, err := yti.OpenFile[TestStruct](
				file.Name(),
				&yti.TableOptions[TestStruct]{MustExist: true},
			)
			assert.NoError(t, err)
		})
	})

	t.Run("open empty file with indices", func(t *testing.T) {
		t.Parallel()
		WithTempFileContent(
			[]TestStruct{
				{
					Id:   1,
					Name: "Test",
				},
			}, func(file *os.File) {
				_, err := yti.OpenFile[TestStruct](
					file.Name(),
					&yti.TableOptions[TestStruct]{
						Indices: TestStructIndices,
					},
				)
				assert.NoError(t, err)
			})
	})
}

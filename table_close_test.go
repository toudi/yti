package yti_test

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/toudi/yti"
)

func TestTableClose(t *testing.T) {
	t.Run("empty table", func(t *testing.T) {
		t.Parallel()
		table, err := yti.OpenFile[TestStruct]("a path", nil)
		assert.NoError(t, err)
		assert.Nil(t, table.Close())
	})

	t.Run("empty file - create directories", func(t *testing.T) {
		t.Parallel()
		dirname := os.TempDir()

		filename := path.Join(dirname, "a", "b", "c", "d.yaml")
		table, err := yti.OpenFile[TestStruct](filename, &yti.TableOptions[TestStruct]{
			MkDirs: true,
		})
		assert.NoError(t, err)
		defer os.RemoveAll(path.Dir(filename))
		assert.NoError(t, table.Insert(TestStruct{Id: 1, Name: "test"}))
		assert.NoError(t, table.Close())
		info, err := os.Stat(filename)
		assert.NoError(t, err)
		assert.Equal(t, false, info.IsDir())
	})

	t.Run("atomic save", func(t *testing.T) {
		t.Parallel()

		tmpTargetFile, err := os.CreateTemp("", "")
		// make sure it's a valid YAML file though with no content
		_, _ = tmpTargetFile.WriteString("---\n")
		assert.NoError(t, err)
		// ok so this is the *target* file which we will want to save to.
		defer os.Remove(tmpTargetFile.Name())
		// now let's use it as input to OpenFile:
		table, err := yti.OpenFile[TestStruct](tmpTargetFile.Name(), nil)
		assert.NoError(t, err)
		// we also have to call Insert, otherwise the file won't save
		// as it won't be marked as dirty
		assert.NoError(t, table.Insert(TestStruct{Id: 1, Name: "test"}))
		table.Close()
		assert.NoError(t, err)
	})
}

package yti_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/toudi/yti"
	"gopkg.in/yaml.v3"
)

type TestStruct struct {
	Id       uint8  `yaml:"id"`
	Name     string `yaml:"name"`
	SomeProp bool   `yaml:"a-prop"`
}

const (
	IndexId = "id"
)

var TestStructIndices = map[string]yti.Indexer[TestStruct]{
	IndexId: func(item TestStruct) interface{} { return item.Id },
}

func WithTempFile(handler func(file *os.File)) {
	file, err := os.CreateTemp("", "")
	if err == nil {
		defer os.Remove(file.Name())
		handler(file)
	}
}

func WithTempFileContent(items []TestStruct, handler func(file *os.File)) {
	WithTempFile(func(file *os.File) {
		err := yaml.NewEncoder(file).Encode(items)
		if err == nil {
			file.Close()
			handler(file)
		}
	})
}

func InMemoryTable(t *testing.T, options *yti.TableOptions[TestStruct]) *yti.Table[TestStruct] {
	var tableOptions = &yti.TableOptions[TestStruct]{
		Indices: TestStructIndices,
	}
	if options != nil {
		tableOptions = options
	}
	table, err := yti.OpenFile[TestStruct]("", tableOptions)
	if err != nil {
		assert.Failf(t, "unexpected error", "err=%v", err)
	}
	return table
}

func InMemoryPopulatedTable(
	t *testing.T,
	values []TestStruct,
	options *yti.TableOptions[TestStruct],
) *yti.Table[TestStruct] {
	var tableOptions = &yti.TableOptions[TestStruct]{
		Indices: TestStructIndices,
	}
	if options != nil {
		tableOptions = options
	}
	table, err := yti.OpenFile[TestStruct]("", tableOptions)
	if err != nil {
		assert.Failf(t, "unexpected error", "err=%v", err)
	}
	for _, value := range values {
		if err := table.Insert(value); err != nil {
			assert.Failf(t, "unexpected error", "err=%v", err)
		}
	}
	return table
}

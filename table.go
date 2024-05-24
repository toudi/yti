package yti

type Item any

type Indexer[V any] func(item V) interface{}

type TableOptions[I Item] struct {
	Indices map[string]Indexer[I]
	// if the file does not exist and this property is set,
	// OpenTable will propagate os.DoesNotExist
	MustExist bool
	// if the library should call MakeDirs when the
	// target file does not exist
	MkDirs bool
}

type Table[I Item] struct {
	options *TableOptions[I]
	items   []I
	indices map[string]Index

	filePath         string
	fileDoesNotExist bool
	dirty            bool
	zeroValue        I
}

func (t *Table[I]) Count() int {
	return len(t.items)
}

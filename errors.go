package yti

import "errors"

var (
	ErrTableAlreadyDefined = errors.New("table already defined")
	ErrEmptyCollection     = errors.New(
		"uninitialized collection. please register some tables first",
	)
	ErrUnknownTable     = errors.New("unknown table")
	ErrLoadingTable     = errors.New("cannot load table")
	ErrItemDoesNotExist = errors.New("cannot find specified item")
	ErrUnknownIndex     = errors.New("unknown index")
	ErrRowsNotUnique    = errors.New("rows with this value are not unique within the table")
)

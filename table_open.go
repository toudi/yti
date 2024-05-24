package yti

import (
	"errors"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

func Open[I Item](input io.Reader, options *TableOptions[I]) (*Table[I], error) {
	table := &Table[I]{
		options: options,
	}

	err := yaml.NewDecoder(input).Decode(&table.items)
	if err != nil {
		return nil, errors.Join(err, ErrLoadingTable)
	}

	if err = table.indexItems(); err != nil {
		return nil, err
	}

	return table, nil
}

func OpenFile[I Item](filePath string, options *TableOptions[I]) (*Table[I], error) {
	file, err := os.Open(filePath)

	// should we error out when the file does not exist ?
	var fileDoesNotExist = os.IsNotExist(err)

	if fileDoesNotExist && options != nil && options.MustExist {
		return nil, err
	}

	// no. let's just treat it as empty
	defer file.Close()

	table := &Table[I]{
		options:          options,
		filePath:         filePath,
		fileDoesNotExist: fileDoesNotExist,
		items:            nil,
	}

	// if the file does not exist don't bother decoding it to slice
	if !fileDoesNotExist {
		err = yaml.NewDecoder(file).Decode(&table.items)
		if err != nil {
			return nil, ErrLoadingTable
		}

		if err = table.indexItems(); err != nil {
			return nil, err
		}
	}

	return table, nil
}

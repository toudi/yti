package yti

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func (t *Table[I]) Close() error {
	// there are no modifications so we don't have to save
	// also, if there's no filePath set it means that a file was not
	// used for opening the table
	if !t.dirty || t.filePath == "" {
		return nil
	}

	if t.fileDoesNotExist && t.options.MkDirs {
		if err := os.MkdirAll(filepath.Dir(t.filePath), 0775); err != nil {
			return err
		}
	}

	if !t.fileDoesNotExist {
		return t.atomicSave()
	}

	file, err := os.OpenFile(
		t.filePath,
		os.O_CREATE|os.O_RDWR,
		0644,
	)

	if err != nil {
		return err
	}

	defer file.Close()

	return yaml.NewEncoder(file).Encode(t.items)
}

func (t *Table[I]) atomicSave() error {
	// because the file already exist, we have to resort to a kind of a hack.
	// let's output the contents to a temporary file and then use the Move
	// function to make sure that it stays atomic:
	tmpFile, err := os.CreateTemp("", "")
	if err != nil {
		return err
	}
	err = yaml.NewEncoder(tmpFile).Encode(t.items)
	if err != nil {
		return err
	}
	if err = tmpFile.Close(); err != nil {
		return err
	}
	return os.Rename(tmpFile.Name(), t.filePath)
}

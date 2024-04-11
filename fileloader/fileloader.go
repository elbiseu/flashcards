package fileloader

import (
	"encoding/csv"
	"os"
)

type FileLoader struct {
	file *os.File
}

func NewFileLoader(name string) (*FileLoader, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}(file)
	return &FileLoader{file: file}, nil
}

func (f *FileLoader) CSV() ([][]string, error) {
	return csv.NewReader(f.file).ReadAll()
}

package external_file

import (
	"os"
)

type TextFile struct {
	Path string
	File *os.File
}

func NewTextFile(path string) (*TextFile, error) {
	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	return &TextFile{Path: path, File: file}, nil
}

func (t *TextFile) CloseFile() error {
	return t.File.Close()
}

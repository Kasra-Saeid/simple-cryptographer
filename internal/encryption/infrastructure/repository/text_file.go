package repository

import (
	"fmt"
	external_file "simple_cryptographer/pkg/external_files"
)

type TextFile struct {
	textFilePkg *external_file.TextFile
}

func NewTextFile(textFilePkg *external_file.TextFile) *TextFile {
	return &TextFile{textFilePkg: textFilePkg}
}

func (t *TextFile) Save(text, encryptedText string) error {
	_, err := t.textFilePkg.File.WriteString(fmt.Sprintf("encryption: %s => %s \n", text, encryptedText))
	if err != nil {
		return err
	}

	return nil
}

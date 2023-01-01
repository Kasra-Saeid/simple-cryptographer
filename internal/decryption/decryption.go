package decryption

import (
	"simple_cryptographer/internal/decryption/domain/port"
	"simple_cryptographer/internal/decryption/domain/service"
	"simple_cryptographer/internal/encryption/infrastructure/repository"
	external_file "simple_cryptographer/pkg/external_files"
)

type Decryption struct {
	DecryptionService service.Decryption
}

func NewDecryption(cryptoPkg port.CryptoPkg, textFilePkg external_file.TextFile) Decryption {
	textFileRepository := repository.NewTextFile(&textFilePkg)
	decryptionService := service.NewDecryption(cryptoPkg, textFileRepository)
	decrypt := Decryption{
		DecryptionService: decryptionService,
	}

	return decrypt
}

func (d Decryption) GetService() service.Decryption {
	return d.DecryptionService
}

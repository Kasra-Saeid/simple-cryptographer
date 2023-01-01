package encryption

import (
	"simple_cryptographer/internal/encryption/domain/port"
	"simple_cryptographer/internal/encryption/domain/service"
	"simple_cryptographer/internal/encryption/infrastructure/repository"
	external_file "simple_cryptographer/pkg/external_files"
)

type Encryption struct {
	encryptionService service.Encryption
}

func NewEncryption(cryptoPkg port.CryptoPkg, textFilePkg external_file.TextFile) Encryption {
	textFileRepository := repository.NewTextFile(&textFilePkg)
	encryptionService := service.NewEncryption(cryptoPkg, textFileRepository)
	encrypt := Encryption{encryptionService: encryptionService}
	return encrypt
}

func (e Encryption) GetService() service.Encryption {
	return e.encryptionService
}

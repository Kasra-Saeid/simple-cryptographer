package service

import (
	"simple_cryptographer/internal/encryption/domain/port"
)

type Encryption struct {
	cryptoPkg port.CryptoPkg
	cacheRepo port.CacheRepo
}

func NewEncryption(cryptoPkg port.CryptoPkg, cacheRepo port.CacheRepo) Encryption {
	return Encryption{cryptoPkg: cryptoPkg, cacheRepo: cacheRepo}
}

func (e Encryption) ToEncryptedText(text string) (string, error) {
	return e.cryptoPkg.Encrypt(text)
}

func (e Encryption) SaveToFile(text, encryptedText string) error {
	return e.cacheRepo.Save(text, encryptedText)
}

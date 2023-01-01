package service

import "simple_cryptographer/internal/decryption/domain/port"

type Decryption struct {
	cryptoPkg port.CryptoPkg
	cacheRepo port.CacheRepo
}

func NewDecryption(cryptoPkg port.CryptoPkg, cacheRepo port.CacheRepo) Decryption {
	return Decryption{cryptoPkg: cryptoPkg, cacheRepo: cacheRepo}
}

func (d Decryption) ToDecryptedText(text string) (string, error) {
	return d.cryptoPkg.Decrypt(text)
}

func (d Decryption) SaveToFile(text, decryptedText string) error {
	return d.cacheRepo.Save(text, decryptedText)
}

package port

type CryptoPkg interface {
	Decrypt(text string) (string, error)
}

type CacheRepo interface {
	Save(text string, encryptedText string) error
}

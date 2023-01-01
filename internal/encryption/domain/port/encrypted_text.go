package port

type CryptoPkg interface {
	Encrypt(text string) (string, error)
}

type CacheRepo interface {
	Save(text string, encryptedText string) error
}

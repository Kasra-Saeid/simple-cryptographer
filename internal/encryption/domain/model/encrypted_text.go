package model

type EncryptedText string

func NewEncryptedText(encryptedText string) EncryptedText {
	return EncryptedText(encryptedText)
}

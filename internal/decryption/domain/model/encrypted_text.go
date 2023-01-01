package model

type DecryptedText string

func NewDecryptedText(decryptedText string) DecryptedText {
	return DecryptedText(decryptedText)
}

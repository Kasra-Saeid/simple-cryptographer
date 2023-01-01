package cryptography

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"log"
)

type Cryptography struct {
	SecretKey string
}

func New(secretKey string) *Cryptography {
	return &Cryptography{
		SecretKey: secretKey,
	}
}

func (c *Cryptography) Encrypt(stringToEncrypt string) (string, error) {
	bytes := []byte("this is bytes go")
	block, err := aes.NewCipher([]byte(c.SecretKey))
	if err != nil {
		return "", err
	}
	plainText := []byte(stringToEncrypt)
	cfb := cipher.NewCFBEncrypter(block, bytes)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	return encode(cipherText), nil
}

func (c *Cryptography) Decrypt(encryptedString string) (string, error) {
	bytes := []byte("this is bytes go")
	block, err := aes.NewCipher([]byte(c.SecretKey))
	if err != nil {
		return "", err
	}

	cipherText := decode(encryptedString)
	cfb := cipher.NewCFBDecrypter(block, bytes)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText), nil

}

func decode(text string) []byte {
	data, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		log.Fatalln(fmt.Errorf("cryptography - decode - DecodeString: %w", err))
	}
	return data
}

func encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

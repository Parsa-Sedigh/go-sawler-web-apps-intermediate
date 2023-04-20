package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

type Encryption struct {
	Key []byte
}

func (e *Encryption) Encrypt(text string) (string, error) {
	plainText := []byte(text)
	block, err := aes.NewCipher(e.Key) // e.Key is our secret key and it must be exactly 32 characters long
	if err != nil {
		return "", err
	}

	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	/* We need to put cipherText on a webpage and we know on webpage we can display text but we can't display certain kinds of things like a slice of bytes([]byte),
	so we need to convert it to base64*/
	return base64.URLEncoding.EncodeToString(cipherText), nil
}

// Decrypt decrypts the cryptoText which is an encrypted text
func (e *Encryption) Decrypt(cryptoText string) (string, error) {
	// BTW you shouldn't ignore this err!
	cipherText, _ := base64.URLEncoding.DecodeString(cryptoText)
	block, err := aes.NewCipher(e.Key) // e.Key must be exactly 32 characters long
	if err != nil {
		return "", err
	}

	if len(cipherText) < aes.BlockSize {
		return "", err
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText) // seems odd that it takes cipherText here twice!

	return fmt.Sprintf("%s", cipherText), nil
}

package server

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
)

func (s *Server) decrypt(data string) string {
	iv := []byte(s.iv)

	block, err := aes.NewCipher([]byte(s.secret))
	if err != nil {
		return ""
	}

	cipherText, err := hex.DecodeString(data)
	if err != nil {
		return ""
	}

	if block.BlockSize() != len(iv) {
		return ""
	}

	ctr := cipher.NewCTR(block, iv)
	plainText := make([]byte, len(cipherText))
	ctr.XORKeyStream(plainText, cipherText)

	return string(plainText)
}

func (s *Server) encrypt(data string) string {
	secret := []byte(s.secret)
	iv := []byte(s.iv)
	plainText := []byte(data)

	block, err := aes.NewCipher(secret)
	if err != nil {
		panic(err)
	}

	cipherText := make([]byte, aes.BlockSize+len(plainText))

	encryptStream := cipher.NewCTR(block, iv)
	encryptStream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	ivHex := hex.EncodeToString(iv)
	encryptedDataHex := hex.EncodeToString(cipherText)

	return encryptedDataHex[len(ivHex):]
}

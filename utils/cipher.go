package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

// AesEncrypt process encrypt with AES alogrithm
func AesEncrypt(data string, key string) string {
	cryptedBytes := _aesEncrypt([]byte(data), []byte(key))
	return string(cryptedBytes[:])
}

func _aesEncrypt(data []byte, key []byte) []byte {
	block, _ := aes.NewCipher(key)
	blockSize := block.BlockSize()
	data = _pKCS7Padding(data, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])

	crypted := make([]byte, len(data))
	blockMode.CryptBlocks(crypted, data)
	return crypted
}

func _pKCS7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padtxt := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtxt...)
}

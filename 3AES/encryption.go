package main

import (
	"aes-encryption/constants"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

func main() {
	plaintext := []byte("123456")
	// 加密 - 使用常量包
	encrypted, err := AESEncrypt1(plaintext, constants.GetAESKey(), constants.GetAESIV())
	if err != nil {
		fmt.Println("Encryption error:", err)
		return
	}
	fmt.Printf("16进制 Encrypted: %x\n", encrypted)
	fmt.Printf("base64编码 Encrypted: %s\n", base64.StdEncoding.EncodeToString(encrypted))

}

func PKCS7Padding1(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}

// AESEncrypt 使用AES进行加密 PKCS7填充
func AESEncrypt1(plaintext, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	plaintext = PKCS7Padding1(plaintext, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv) //CBC分组模式和初始化向量（IV）
	ciphertext := make([]byte, len(plaintext))
	blockMode.CryptBlocks(ciphertext, plaintext)
	return ciphertext, nil
}

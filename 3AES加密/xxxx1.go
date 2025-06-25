package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

func main() {
	var KEY = []byte("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab") //密钥长度必须为 16/24/32 字节。我这里使用的是AES - 256 使用 32 字节密钥
	var IV = []byte("aaaaaaaaaaaaaaax")                  //初始化向量（CBC/CFB/OFB模式需要，长度必须等于块大小16字节）
	plaintext := []byte("123456")
	// 加密
	encrypted, err := AESEncrypt1(plaintext, KEY, IV)
	if err != nil {
		fmt.Println("Encryption error:", err)
		return
	}
	fmt.Printf("Encrypted: %x\n", encrypted)

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

package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
)

// PKCS7Padding 实现PKCS7填充 (函数用于对明文进行填充，以满足 AES 加密块大小的要求)
func PKCS7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}

// PKCS7UnPadding 实现PKCS7去填充
func PKCS7UnPadding(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}

// AESEncrypt 使用AES进行加密 PKCS7
func AESEncrypt(plaintext, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	plaintext = PKCS7Padding(plaintext, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	ciphertext := make([]byte, len(plaintext))
	blockMode.CryptBlocks(ciphertext, plaintext)
	return ciphertext, nil
}

// AESDecrypt 使用AES进行解密
func AESDecrypt(ciphertext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	plaintext := make([]byte, len(ciphertext))
	blockMode.CryptBlocks(plaintext, ciphertext)
	plaintext = PKCS7UnPadding(plaintext)
	return plaintext, nil
}
func main() {
	key := []byte("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa") //32
	iv := []byte("aaaaaaaaaaaaaaaa")                  //16
	plaintext := []byte("123456")
	// 加密
	encrypted, err := AESEncrypt(plaintext, key, iv)
	if err != nil {
		fmt.Println("Encryption error:", err)
		return
	}
	fmt.Printf("Encrypted: %x\n", encrypted) //encrypted 加密过后是16进制字符串(979a78eb039969783473a9caac3afc61)
	//fmt.Println("转换=" + (hexBytesToDecimalStr(encrypted)))

	//[]byte("979a78eb039969783473a9caac3afc61")
	//hexStr := fmt.Sprintf("%x", encrypted) //将加密后的字节数组转换为十六进制字符串。以十六进制形式输出，字母小写。
	// 将十六进制字符串转换回字节数组
	//mybyte, err := hex.DecodeString(hexStr)
	mybyte, err := hex.DecodeString("979a78eb039969783473a9caac3afc61") //直接加双引号也是16进制字符串
	if err != nil {
		fmt.Println("Hex decode error:", err)
		return
	}
	fmt.Printf("mybyte: %x\n", mybyte) //encrypted 加密过后是16进制字节数组
	// 解密
	//decrypted, err := AESDecrypt(encrypted, key)
	decrypted, err := AESDecrypt(mybyte, key)
	if err != nil {
		fmt.Println("Decryption error:", err)
		return
	}
	fmt.Printf("Decrypted: %s\n", decrypted)
}

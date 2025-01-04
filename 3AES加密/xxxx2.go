package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"runtime"
)

func main() {
	var KEY = []byte("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab") //AES - 256 使用 32 字节密钥
	var IV = []byte("aaaaaaaaaaaaaaaa")

	mybyte, err := hex.DecodeString("2635c485dc84a03e7b3c636ca350e6ac") //直接加双引号也是16进制字符串, 转为16进制字节数组
	if err != nil {
		fmt.Println("Hex decode error:", err)
		return
	}

	// 解密
	decrypted, err := AESDecrypt1(mybyte, KEY, IV)
	if err != nil {
		fmt.Println("Decryption error:", err)
		return
	}
	fmt.Printf("Decrypted: %s\n", decrypted)
}

// PKCS7UnPadding1 实现PKCS7去填充
func PKCS7UnPadding1(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, fmt.Errorf("数据为空，无法去填充")
	}
	unpadding := int(data[length-1])
	if unpadding > length {
		return nil, fmt.Errorf("填充数据错误，去填充长度大于数据长度")
	}
	return data[:(length - unpadding)], nil
}

// AESDecrypt1 使用AES进行解密
func AESDecrypt1(ciphertext, key []byte, iv []byte) ([]byte, error) {
	if len(ciphertext) == 0 {
		return nil, fmt.Errorf("密文为空")
	}
	if len(key) == 0 {
		return nil, fmt.Errorf("密钥为空")
	}
	if len(iv) == 0 {
		return nil, fmt.Errorf("初始向量为空")
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("创建AES密码块失败: %w", err)
	}
	blockSize := block.BlockSize()
	if len(ciphertext)%blockSize != 0 {
		return nil, fmt.Errorf("密文长度不是块大小的整数倍")
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	plaintext := make([]byte, len(ciphertext))
	// 捕获CryptBlocks可能出现的错误
	err = func() error {
		defer func() {
			if r := recover(); r != nil {
				if _, ok := r.(runtime.Error); ok {
					panic(r)
				}
			}
		}()
		blockMode.CryptBlocks(plaintext, ciphertext)
		return nil
	}()
	if err != nil {
		return nil, fmt.Errorf("解密过程中发生错误: %w", err)
	}
	unPaddedPlaintext, err := PKCS7UnPadding1(plaintext)
	if err != nil {
		return nil, fmt.Errorf("去填充过程中发生错误: %w", err)
	}
	return unPaddedPlaintext, nil
}

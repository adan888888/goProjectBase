package main

import (
	"aes-encryption/constants"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var plaintextInput = flag.String("text", "", "要加密的明文文本")
	var inputFile = flag.String("file", "", "从文件读取明文")
	var outputFile = flag.String("output", "", "将密文保存到文件（可选）")
	flag.Parse()

	var plaintext []byte
	var err error

	// 优先从文件读取
	if *inputFile != "" {
		plaintext, err = os.ReadFile(*inputFile)
		if err != nil {
			fmt.Printf("❌ 读取文件失败: %v\n", err)
			return
		}
	} else if *plaintextInput != "" {
		plaintext = []byte(*plaintextInput)
	} else {
		// 默认明文
		plaintext = []byte(`{"remarks":"只有usdt哈哈哈哈",}`)
	}

	// 显示明文信息（用于调试）
	plaintextHash := sha256.Sum256(plaintext)
	fmt.Printf("📝 明文长度: %d 字节\n", len(plaintext))
	fmt.Printf("🔐 明文SHA256: %x\n", plaintextHash)
	fmt.Println()

	// 加密 - 使用常量包
	encrypted, err := AESEncrypt1(plaintext, constants.GetAESKey(), constants.GetAESIV())
	if err != nil {
		fmt.Printf("❌ 加密错误: %v\n", err)
		return
	}

	// 计算密文校验和
	encryptedHash := sha256.Sum256(encrypted)
	hexEncrypted := hex.EncodeToString(encrypted)
	base64Encrypted := base64.StdEncoding.EncodeToString(encrypted)

	// 输出到控制台
	fmt.Println("=" + strings.Repeat("=", 70) + "=")
	fmt.Println("🔒 加密结果")
	fmt.Println("=" + strings.Repeat("=", 70) + "=")
	fmt.Printf("📏 密文长度: %d 字节\n", len(encrypted))
	fmt.Printf("🔐 密文SHA256: %x\n", encryptedHash)
	fmt.Println()
	fmt.Printf("16进制 Encrypted: %s\n", hexEncrypted)
	fmt.Println()
	fmt.Printf("base64编码 Encrypted: %s\n", base64Encrypted)
	fmt.Println()
	fmt.Println("=" + strings.Repeat("=", 70) + "=")
	fmt.Println()
	fmt.Println("⚠️  注意: 在 GoLand 终端中复制 16 进制字符串可能丢失字符")
	fmt.Println("💡 推荐使用以下方式:")
	fmt.Println("   1. 使用 base64 格式（更短，不容易出错）")
	fmt.Println("   2. 使用文件方式（完全避免复制错误）:")
	fmt.Println("      go run encryption/main.go -output encrypted.txt")
	fmt.Println("      go run decrypt/main.go -file encrypted.txt")
	fmt.Println()

	// 如果指定了输出文件，保存密文
	if *outputFile != "" {
		output := fmt.Sprintf("HEX: %s\nBASE64: %s\n", hexEncrypted, base64Encrypted)
		err = os.WriteFile(*outputFile, []byte(output), 0644)
		if err != nil {
			fmt.Printf("❌ 保存文件失败: %v\n", err)
		} else {
			fmt.Printf("✅ 密文已保存到: %s\n", *outputFile)

			// 验证文件内容是否正确
			fileContent, readErr := os.ReadFile(*outputFile)
			if readErr == nil {
				// 提取文件中的 HEX 值
				fileLines := strings.Split(string(fileContent), "\n")
				var fileHex string
				for _, line := range fileLines {
					if strings.HasPrefix(line, "HEX: ") {
						fileHex = strings.TrimPrefix(line, "HEX: ")
						break
					}
				}

				// 对比控制台输出和文件内容
				if fileHex != hexEncrypted {
					fmt.Printf("⚠️  警告: 文件中的 HEX 值与控制台输出不一致！\n")
					fmt.Printf("   控制台 HEX 长度: %d 字符\n", len(hexEncrypted))
					fmt.Printf("   文件 HEX 长度: %d 字符\n", len(fileHex))
					fmt.Printf("   差异位置: ")
					minLen := len(hexEncrypted)
					if len(fileHex) < minLen {
						minLen = len(fileHex)
					}
					for i := 0; i < minLen; i++ {
						if hexEncrypted[i] != fileHex[i] {
							start := i - 10
							if start < 0 {
								start = 0
							}
							end := i + 10
							if end > minLen {
								end = minLen
							}
							fmt.Printf("\n   位置 %d: 控制台='%s', 文件='%s'\n",
								i, hexEncrypted[start:end], fileHex[start:end])
							break
						}
					}
					if len(hexEncrypted) != len(fileHex) {
						fmt.Printf("   长度不同: 控制台=%d, 文件=%d\n", len(hexEncrypted), len(fileHex))
					}
					fmt.Printf("\n💡 如果看到不一致，请确保：\n")
					fmt.Printf("   1. 控制台输出和文件是同一运行的结果\n")
					fmt.Printf("   2. 不要从控制台复制，直接使用文件\n")
					fmt.Printf("   3. 使用 base64 格式更可靠\n")
				} else {
					fmt.Printf("✅ 验证通过: 文件内容与控制台输出一致（%d 字符）\n", len(hexEncrypted))
					fmt.Printf("📄 文件中的 HEX: %s\n", fileHex)
				}
			}
		}
	}

	// 提供快速复制提示
	fmt.Println()
	fmt.Println("💡 快速解密（复制以下命令）:")
	fmt.Printf("   go run decrypt/main.go -base64 \"%s\"\n", base64Encrypted)
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

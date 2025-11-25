package main

import (
	"aes-encryption/constants"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"regexp"
	"runtime"
	"strings"
)

// cleanBase64String 清理 base64 字符串，去除所有空白字符和不可见字符
func cleanBase64String(s string) string {
	// 去除所有空白字符（空格、制表符、换行符等）
	re := regexp.MustCompile(`\s+`)
	s = re.ReplaceAllString(s, "")
	// 去除首尾空白
	s = strings.TrimSpace(s)
	return s
}

// cleanHexString 清理 16 进制字符串，去除所有空白字符
func cleanHexString(s string) string {
	// 去除所有空白字符
	re := regexp.MustCompile(`\s+`)
	s = re.ReplaceAllString(s, "")
	// 去除首尾空白
	s = strings.TrimSpace(s)
	return s
}

// extractBase64FromOutput 从加密程序的输出中提取 base64 密文
func extractBase64FromOutput(output string) string {
	// 查找 "base64编码 Encrypted: " 后面的内容
	re := regexp.MustCompile(`base64编码\s+Encrypted:\s*([A-Za-z0-9+/=]+)`)
	matches := re.FindStringSubmatch(output)
	if len(matches) > 1 {
		return matches[1]
	}
	// 查找 "BASE64: " 后面的内容（文件格式）
	re = regexp.MustCompile(`BASE64:\s*([A-Za-z0-9+/=]+)`)
	matches = re.FindStringSubmatch(output)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

// extractHexFromOutput 从加密程序的输出中提取 16 进制密文
func extractHexFromOutput(output string) string {
	// 查找 "16进制 Encrypted: " 后面的内容
	re := regexp.MustCompile(`16进制\s+Encrypted:\s*([0-9a-fA-F]+)`)
	matches := re.FindStringSubmatch(output)
	if len(matches) > 1 {
		return matches[1]
	}
	// 查找 "HEX: " 后面的内容（文件格式）
	re = regexp.MustCompile(`HEX:\s*([0-9a-fA-F]+)`)
	matches = re.FindStringSubmatch(output)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

func main() {
	// 支持命令行参数
	var ciphertextHex = flag.String("hex", "", "16进制格式的密文")
	var ciphertextBase64 = flag.String("base64", "", "base64格式的密文")
	var inputFile = flag.String("file", "", "从文件读取密文（自动识别格式）")
	flag.Parse()

	var ciphertext []byte
	var err error
	var format string

	// 优先从文件读取
	if *inputFile != "" {
		fileContent, err := os.ReadFile(*inputFile)
		if err != nil {
			fmt.Printf("❌ 读取文件失败: %v\n", err)
			return
		}
		content := string(fileContent)

		// 尝试从文件内容中提取 base64
		extractedBase64 := extractBase64FromOutput(content)
		if extractedBase64 != "" {
			cleanedBase64 := cleanBase64String(extractedBase64)
			ciphertext, err = base64.StdEncoding.DecodeString(cleanedBase64)
			if err == nil {
				format = "base64（从文件提取）"
			}
		}

		// 如果提取 base64 失败，尝试提取 16 进制
		if err != nil || extractedBase64 == "" {
			extractedHex := extractHexFromOutput(content)
			if extractedHex != "" {
				cleanedHex := cleanHexString(extractedHex)
				ciphertext, err = hex.DecodeString(cleanedHex)
				if err == nil {
					format = "16进制（从文件提取）"
				}
			}
		}

		// 如果提取失败，尝试直接解析（可能是纯密文文件）
		if err != nil {
			// 尝试作为 base64 解析
			cleanedContent := cleanBase64String(content)
			ciphertext, err = base64.StdEncoding.DecodeString(cleanedContent)
			if err == nil {
				format = "base64（从文件）"
			} else {
				// 尝试作为 16 进制解析
				cleanedContent = cleanHexString(content)
				ciphertext, err = hex.DecodeString(cleanedContent)
				if err == nil {
					format = "16进制（从文件）"
				}
			}
		}

		if err != nil {
			fmt.Printf("❌ 无法从文件中识别密文格式: %v\n", err)
			return
		}
	} else if *ciphertextHex != "" {
		cleanedHex := cleanHexString(*ciphertextHex)
		originalLen := len(*ciphertextHex)
		cleanedLen := len(cleanedHex)

		// 检查长度是否为偶数
		if cleanedLen%2 != 0 {
			fmt.Printf("❌ 16进制字符串长度错误\n")
			fmt.Printf("   原始密文长度: %d 字符\n", originalLen)
			fmt.Printf("   清理后长度: %d 字符（必须是偶数）\n", cleanedLen)
			fmt.Printf("   缺少: %d 字符\n", 1)
			fmt.Printf("\n💡 建议:\n")
			fmt.Printf("   1. 使用 base64 格式（更短，不容易出错）:\n")
			fmt.Printf("      go run decrypt/main.go -base64 \"...\"\n")
			fmt.Printf("   2. 使用文件方式（推荐，完全避免复制错误）:\n")
			fmt.Printf("      go run encryption/main.go -output encrypted.txt\n")
			fmt.Printf("      go run decrypt/main.go -file encrypted.txt\n")
			fmt.Printf("   3. 检查复制的 16 进制字符串是否完整\n")
			return
		}

		ciphertext, err = hex.DecodeString(cleanedHex)
		format = "16进制"
		if err != nil {
			fmt.Printf("❌ 16进制解码错误: %v\n", err)
			fmt.Printf("   原始密文长度: %d 字符\n", originalLen)
			fmt.Printf("   清理后长度: %d 字符\n", cleanedLen)
			previewLen := 50
			if cleanedLen < previewLen {
				previewLen = cleanedLen
			}
			fmt.Printf("   清理后的密文（前%d字符）: %s", previewLen, cleanedHex[:previewLen])
			if cleanedLen > previewLen {
				fmt.Printf("...")
			}
			fmt.Println()
			if cleanedLen > previewLen {
				start := cleanedLen - previewLen
				if start < 0 {
					start = 0
				}
				fmt.Printf("   清理后的密文（后%d字符）: ...%s\n", previewLen, cleanedHex[start:])
			}
			fmt.Printf("\n💡 建议:\n")
			fmt.Printf("   1. 使用 base64 格式（更短，不容易出错）:\n")
			fmt.Printf("      go run decrypt/main.go -base64 \"...\"\n")
			fmt.Printf("   2. 使用文件方式（推荐，完全避免复制错误）:\n")
			fmt.Printf("      go run encryption/main.go -output encrypted.txt\n")
			fmt.Printf("      go run decrypt/main.go -file encrypted.txt\n")
			return
		}
	} else if *ciphertextBase64 != "" {
		cleanedBase64 := cleanBase64String(*ciphertextBase64)
		ciphertext, err = base64.StdEncoding.DecodeString(cleanedBase64)
		format = "base64"
		if err != nil {
			fmt.Printf("❌ base64解码错误: %v\n", err)
			fmt.Printf("   原始密文长度: %d 字符\n", len(*ciphertextBase64))
			fmt.Printf("   清理后长度: %d 字符\n", len(cleanedBase64))
			fmt.Printf("   清理后的密文: %s\n", cleanedBase64)
			fmt.Printf("   提示: 请检查密文是否完整，是否包含特殊字符\n")
			return
		}
	} else {
		// 如果没有命令行参数，尝试从标准输入读取
		if len(os.Args) > 1 {
			input := strings.Join(os.Args[1:], " ")

			// 首先尝试从加密程序输出中提取 base64
			extractedBase64 := extractBase64FromOutput(input)
			extractedHex := extractHexFromOutput(input)
			success := false

			if extractedBase64 != "" {
				cleanedBase64 := cleanBase64String(extractedBase64)
				ciphertext, err = base64.StdEncoding.DecodeString(cleanedBase64)
				if err == nil {
					format = "base64（从输出中提取）"
					success = true
				}
			}

			// 如果提取 base64 失败，尝试提取 16 进制
			if !success && extractedHex != "" {
				cleanedHex := cleanHexString(extractedHex)
				ciphertext, err = hex.DecodeString(cleanedHex)
				if err == nil {
					format = "16进制（从输出中提取）"
					success = true
				}
			}

			// 如果提取失败，尝试直接解析
			if !success {
				// 尝试作为16进制解析
				cleanedInput := cleanHexString(input)
				ciphertext, err = hex.DecodeString(cleanedInput)
				if err != nil {
					// 如果16进制解析失败，尝试base64
					cleanedInput = cleanBase64String(input)
					ciphertext, err = base64.StdEncoding.DecodeString(cleanedInput)
					if err != nil {
						fmt.Println("❌ 无法识别密文格式")
						fmt.Println("\n使用方法:")
						fmt.Println("  方法1: 从文件读取（推荐，避免复制错误）")
						fmt.Println("    go run decrypt/main.go -file encrypted.txt")
						fmt.Println("\n  方法2: 使用参数指定格式")
						fmt.Println("    go run decrypt/main.go -hex <16进制密文>")
						fmt.Println("    go run decrypt/main.go -base64 <base64密文>")
						fmt.Println("\n  方法3: 直接粘贴加密程序的完整输出（自动提取）")
						fmt.Println("    go run decrypt/main.go 'base64编码 Encrypted: X6UqEhXn8Hye0gBG...'")
						fmt.Println("\n  方法4: 直接传入密文（自动识别格式）")
						fmt.Println("    go run decrypt/main.go X6UqEhXn8Hye0gBG...")
						return
					}
					format = "base64"
				} else {
					format = "16进制"
				}
			}
		} else {
			fmt.Println("❌ 请提供密文")
			fmt.Println("\n使用方法:")
			fmt.Println("  方法1: 从文件读取（推荐，避免复制错误）")
			fmt.Println("    go run decrypt/main.go -file encrypted.txt")
			fmt.Println("\n  方法2: 使用参数指定格式")
			fmt.Println("    go run decrypt/main.go -hex <16进制密文>")
			fmt.Println("    go run decrypt/main.go -base64 <base64密文>")
			fmt.Println("\n  方法3: 直接粘贴加密程序的完整输出（自动提取）")
			fmt.Println("    go run decrypt/main.go 'base64编码 Encrypted: X6UqEhXn8Hye0gBG...'")
			fmt.Println("\n  方法4: 直接传入密文（自动识别格式）")
			fmt.Println("    go run decrypt/main.go X6UqEhXn8Hye0gBG...")
			return
		}
	}

	// 显示密文信息
	block, _ := aes.NewCipher(constants.GetAESKey())
	blockSize := block.BlockSize()
	fmt.Printf("📝 密文格式: %s\n", format)
	fmt.Printf("📏 密文字节长度: %d 字节\n", len(ciphertext))
	fmt.Printf("📦 AES块大小: %d 字节\n", blockSize)
	if len(ciphertext)%blockSize != 0 {
		fmt.Printf("❌ 错误: 密文长度 %d 不是块大小 %d 的整数倍\n", len(ciphertext), blockSize)
		fmt.Printf("   缺少 %d 字节\n", blockSize-(len(ciphertext)%blockSize))
		if format == "16进制" {
			fmt.Printf("\n⚠️  16进制字符串在 GoLand 终端中复制时可能丢失字符\n")
			fmt.Printf("\n💡 建议:\n")
			fmt.Printf("   1. 使用 base64 格式（更短，不容易出错）:\n")
			fmt.Printf("      go run decrypt/main.go -base64 \"...\"\n")
			fmt.Printf("   2. 使用文件方式（推荐，完全避免复制错误）:\n")
			fmt.Printf("      go run encryption/main.go -output encrypted.txt\n")
			fmt.Printf("      go run decrypt/main.go -file encrypted.txt\n")
			fmt.Printf("   3. 检查复制的 16 进制字符串是否完整\n")
		}
		return
	}

	// 解密 - 使用常量包
	decrypted, err := AESDecrypt1(ciphertext, constants.GetAESKey(), constants.GetAESIV())
	if err != nil {
		fmt.Printf("❌ 解密错误: %v\n", err)
		return
	}
	fmt.Printf("✅ 解密成功: %s\n", decrypted)
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
		return nil, fmt.Errorf("密文长度 %d 不是块大小 %d 的整数倍（缺少 %d 字节）",
			len(ciphertext), blockSize, blockSize-(len(ciphertext)%blockSize))
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

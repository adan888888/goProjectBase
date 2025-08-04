package main

import (
	"crypto/rand"
	"fmt"
)

func main() {
	var num1 int8 = 10
	fmt.Println(num1)

	var s1 = []string{"dfdfdfdf"}
	fmt.Println(s1)
	//通过 []byte(字符串) 的转换，会将字符串中的每个字符转换为对应的 ASCII 字节值，存储到切片中。
	var b1 = []byte("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab")
	fmt.Println(len(b1))
	fmt.Println(b1)
	fmt.Println("字符形式：", string(b1)) // 转换为字符串后打印（显示字符）

	randomBytes, _ := generatePrintableRandomBytes(32)
	fmt.Println("随机生成的是", randomBytes)
	fmt.Println("字符形式", string(randomBytes))

}

// 生成仅包含可打印字符的随机字节（32~126范围）
func generatePrintableRandomBytes(length int) ([]byte, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	// 调整每个字节到32~126（可打印ASCII范围）
	for i := range b {
		b[i] = 32 + (b[i] % 95) // 95 = 126 - 32 + 1
	}
	return b, nil
}

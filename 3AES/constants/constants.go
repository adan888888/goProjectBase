package constants

// AES 加密相关常量
const (
	// AES_KEY AES-256 密钥（32字节）
	AES_KEY = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	
	// AES_IV 初始化向量（16字节）
	AES_IV = "0123456789abcdef"
)

// 获取 AES 密钥字节数组
func GetAESKey() []byte {
	return []byte(AES_KEY)
}

// 获取 AES 初始化向量字节数组
func GetAESIV() []byte {
	return []byte(AES_IV)
}

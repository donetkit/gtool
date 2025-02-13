package ghash

import (
	"crypto/sha256"
	"encoding/hex"
)

func Sha256(input []byte) string {
	// 创建一个新的 SHA-256 哈希接口
	data := sha256.New()

	// 将输入字符串写入哈希接口中
	data.Write(input)

	// 计算哈希值，并返回其字节切片表示形式
	hashBytes := data.Sum(nil)

	// 将字节切片转换为十六进制字符串表示形式
	return hex.EncodeToString(hashBytes)
}

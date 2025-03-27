package tool

import (
	"crypto/md5"
	"fmt"
	"io"
)

// 对s进行md5加密
// 详细说明见：https://pkg.go.dev/crypto/md5
func GenerateMD5(s string) string {
	// 新建一个用于md5加密的对象
	h := md5.New()
	// 写入s
	io.WriteString(h, s)
	// 导出
	return fmt.Sprintf("%x", h.Sum(nil))
}

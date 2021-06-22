package helper

import (
	"crypto/sha1"
	"io"
	"os"
	"path/filepath"
)

// CheckSum 用于计算文件的校验码
func CheckSum(file string) ([]byte, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	sha1 := sha1.New()
	io.Copy(sha1, f)
	return sha1.Sum(nil), nil
}

// GetBaseName 获取文件名
func GetBaseName(path string) string {
	return filepath.Base(path)
}

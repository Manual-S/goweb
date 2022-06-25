package util

import "os"

// GetExecDirectory 获取当前的运行目录
func GetExecDirectory() string {
	file, err := os.Getwd()
	if err != nil {
		return ""
	}

	return file + "/"
}

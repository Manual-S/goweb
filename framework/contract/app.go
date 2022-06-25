// Package contract 应用目录服务
package contract

const DirectoryKey = "DirectoryKey"

// DirectoryInf 目录服务的接口协议
type DirectoryInf interface {
	Version() string

	BaseFolder() string

	// ConfigFolder 获取配置文件的路径
	ConfigFolder() string
}

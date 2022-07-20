// Package contract 应用目录服务
package contract

const DirectoryKey = "DirectoryKey"

// DirectoryInf 目录服务的接口协议
type DirectoryInf interface {
	Version() string

	BaseFolder() string

	// ConfigFolder 获取配置文件的路径
	ConfigFolder() string

	// RuntimeFolder 定义业务的运行中间态信息
	RuntimeFolder() string

	// LogFolder 日志存放的地址
	LogFolder() string

	// AppID 获取当前服务的appid
	AppID() string
}

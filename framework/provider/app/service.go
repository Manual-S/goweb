// Package app 服务的具体实现
package app

import (
	"errors"
	"flag"
	"github.com/google/uuid"
	"goweb/framework"
	"goweb/framework/contract"
	"goweb/framework/util"
	"path/filepath"
)

type Directory struct {
	contract.DirectoryInf

	container framework.Container
	// 基础路径
	baseFolder string
	// 表示当前服务的唯一id 可以用于分布式锁
	appID string
}

func NewDirectoryService(params ...interface{}) (interface{}, error) {
	if len(params) < 2 {
		return nil, errors.New("params error")
	}
	container := params[0].(framework.Container)
	baseFolder := params[1].(string)
	appID := uuid.New().String()
	return &Directory{
		container:  container,
		baseFolder: baseFolder,
		appID:      appID,
	}, nil
}

func (d *Directory) Version() string {
	return ""
}

func (d *Directory) BaseFolder() string {
	if d.baseFolder != "" {
		return d.baseFolder
	}

	// 如果没有设置 就使用参数
	var baseFolder string
	flag.StringVar(&baseFolder, "base_folder", "", "base_folder的参数")
	flag.Parse()

	if baseFolder != "" {
		return baseFolder
	}

	return util.GetExecDirectory()
}

func (d *Directory) StorageFolder() string {
	return filepath.Join(d.BaseFolder(), "storage")
}

// ConfigFolder 获取配置文件的路径
func (d *Directory) ConfigFolder() string {
	return ""
}

// RuntimeFolder 定义业务运行的中间态信息
func (d *Directory) RuntimeFolder() string {
	return filepath.Join(d.StorageFolder(), "runtime")
}

// LogFolder 定义日志存储的信息
func (d *Directory) LogFolder() string {
	return filepath.Join(d.StorageFolder(), "log")
}

func (d *Directory) AppID() string {
	return d.appID
}

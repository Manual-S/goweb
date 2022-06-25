// Package app 服务的具体实现
package app

import (
	"errors"
	"flag"
	"goweb/framework"
	"goweb/framework/contract"
	"goweb/framework/util"
)

type Directory struct {
	contract.DirectoryInf

	container framework.Container
	// 基础路径
	baseFolder string
}

func NewDirectoryService(params ...interface{}) (interface{}, error) {
	if len(params) < 2 {
		return nil, errors.New("params error")
	}
	container := params[0].(framework.Container)
	baseFolder := params[1].(string)

	return &Directory{
		container:  container,
		baseFolder: baseFolder,
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

// ConfigFolder 获取配置文件的路径
func (d *Directory) ConfigFolder() string {
	return ""
}

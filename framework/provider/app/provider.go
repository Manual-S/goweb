// Package app 服务的提供者
package app

import (
	"goweb/framework"
	"goweb/framework/contract"
)

// DirectoryProvider 服务提供者
type DirectoryProvider struct {
	framework.ServiceProvider

	BaseFolder string
}

// Name 获取服务提供者的名字
func (d *DirectoryProvider) Name() string {
	return contract.DirectoryKey
}

// IsDefer 决定是否要延迟初始化
func (d *DirectoryProvider) IsDefer() bool {
	return false
}

// Boot 调用实例化服务时调用
func (d *DirectoryProvider) Boot(container framework.Container) error {
	return nil
}

// Register 初始化一个服务
func (d *DirectoryProvider) Register(container framework.Container) framework.NewInstance {
	return NewDirectoryService
}

// Params 获取params参数
func (d *DirectoryProvider) Params(container framework.Container) []interface{} {
	return []interface{}{container, d.BaseFolder}
}

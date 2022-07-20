package distributed

import (
	"goweb/framework"
	"goweb/framework/contract"
)

type LocalDistributeProvider struct {
	framework.ServiceProvider
}

// Name 获取服务提供者的名字
func (l *LocalDistributeProvider) Name() string {
	return contract.DistributedKey
}

// IsDefer 决定是否要延迟初始化
func (l *LocalDistributeProvider) IsDefer() bool {
	return false
}

// Boot 调用实例化服务时调用
func (l *LocalDistributeProvider) Boot(container framework.Container) error {
	return nil
}

// Register 初始化一个服务
func (l *LocalDistributeProvider) Register(container framework.Container) framework.NewInstance {
	return NewLocalDistributedService
}

// Params 获取params参数
func (l *LocalDistributeProvider) Params(container framework.Container) []interface{} {
	return []interface{}{container}
}

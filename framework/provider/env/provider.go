package env

import (
	"goweb/framework"
	"goweb/framework/contract"
)

type EnvProvider struct {
	framework.ServiceProvider
}

// Name 获取服务提供者的名字
func (e *EnvProvider) Name() string {
	return contract.EnvKey
}

// IsDefer 决定是否要延迟初始化
func (e *EnvProvider) IsDefer() bool {
	return false
}

// Boot 调用实例化服务时调用
func (e *EnvProvider) Boot(container framework.Container) error {
	return nil
}

// Register 初始化一个服务
func (e *EnvProvider) Register(container framework.Container) framework.NewInstance {
	return NewEnvService
}

// Params 获取params参数
func (e *EnvProvider) Params(container framework.Container) []interface{} {
	return nil
}

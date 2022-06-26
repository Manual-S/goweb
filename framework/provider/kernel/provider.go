package kernel

import (
	"goweb/framework"
	"goweb/framework/contract"
	"goweb/framework/gin"
)

// KernelProvider 服务提供者
type KernelProvider struct {
	framework.ServiceProvider

	HttpEngine *gin.Engine
}

// Name 获取服务提供者的名字
func (k *KernelProvider) Name() string {
	return contract.KernelKey
}

// IsDefer 决定是否要延迟初始化
func (k *KernelProvider) IsDefer() bool {
	return false
}

// Boot 调用实例化服务时调用
func (k *KernelProvider) Boot(container framework.Container) error {
	if k.HttpEngine != nil {
		k.HttpEngine = gin.Default()
	}
	k.HttpEngine.SetContainer(container)

	return nil
}

// Register 初始化一个服务
func (k *KernelProvider) Register(container framework.Container) framework.NewInstance {
	return NewKernelService
}

// Params 获取params参数
func (k *KernelProvider) Params(container framework.Container) []interface{} {
	return []interface{}{k.HttpEngine}
}

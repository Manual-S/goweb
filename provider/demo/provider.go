package demo

import "goweb/framework"

type DemoServiceProvider struct {
	framework.ServiceProvider
}

// Name 获取服务提供者的名字
func (d *DemoServiceProvider) Name() string {
	return key
}

// IsDefer 决定是否要延迟初始化
func (d *DemoServiceProvider) IsDefer() bool {
	return true
}

// Boot 调用实例化服务时调用
func (d *DemoServiceProvider) Boot(container framework.Container) error {
	// Boot方法什么都不做
	return nil
}

// Register 初始化一个服务
func (d *DemoServiceProvider) Register(container framework.Container) framework.NewInstance {
	return NewDemoService
}

// Params 获取params参数
func (d *DemoServiceProvider) Params(container framework.Container) []interface{} {
	return []interface{}{container}
}

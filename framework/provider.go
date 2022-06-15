package framework

type NewInstance func(...interface{}) (interface{}, error)

// ServiceProvider 服务提供者
type ServiceProvider interface {
	// Name 获取服务提供者的名字
	Name() string
	// IsDefer 决定是否要延迟初始化
	IsDefer() bool
	// Boot 调用实例化服务时调用
	Boot(container Container) error
	// Register 初始化一个服务
	Register(container Container) NewInstance
	// Params 获取params参数
	Params(container Container) []interface{}
}

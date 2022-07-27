package contract

const (
	EnvProduction  = "product"
	EnvTest        = "test"
	EnvDevelopment = "development"

	EnvKey = "envkey"
)

type EnvInf interface {
	// AppEnv 获取当前的环境
	AppEnv() string
	// IsExist  判断一个环境变量是否被设置
	IsExist(string) bool
	// Get 获取某个环境变量 如果没有设置就返回""
	Get(string) string

	// All 获取所有的环境变量
	All() map[string]string
}

package http

import "goweb/framework/gin"

// NewHttpEngine 返回一个绑定了路由的web引擎
func NewHttpEngine() (*gin.Engine, error) {
	// 这里如果不设置为调试模式
	// 在启动的时候会输出调试信息
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// 注册路由
	RegisterRouter(r)

	return r, nil
}

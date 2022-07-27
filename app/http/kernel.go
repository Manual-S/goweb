package http

import "goweb/framework/gin"

// NewHttpEngine 返回一个绑定了路由的web引擎
func NewHttpEngine() (*gin.Engine, error) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// 注册路由
	RegisterRouter(r)

	return r, nil
}

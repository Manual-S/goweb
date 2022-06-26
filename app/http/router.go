package http

import (
	"goweb/app/http/module/demo"
	"goweb/framework/gin"
)

func RegisterRouter(r *gin.Engine) {
	// 注册路由
	demo.Register(r)
}
